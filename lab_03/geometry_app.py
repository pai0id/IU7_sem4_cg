import tkinter as tk  # Импорт модуля для создания графического интерфейса
from tkinter import ttk
from tkinter import colorchooser
from tkinter import messagebox
from math import pi, sin, cos, radians, floor, fabs
import numpy as np
import matplotlib.pyplot as plt
import time
import colorutils as cu

class GeometryApp:

    def set_color(self):
        self.color_code = colorchooser.askcolor(title="Choose color")[0]
        
    def set_bg_color(self):
        self.bg_color_code = colorchooser.askcolor(title="Choose color")[1]
        self.canvas.configure(bg=self.bg_color_code)

    def parse_range(self):
        try:
            line_len = float(self.len_val.get())
            angle_spin = float(self.ang_val.get())
        except:
            messagebox.showerror("Ошибка", "Неверно введены координаты")
            return

        if (line_len <= 0):
            messagebox.showerror("Ошибка", "Длина линии должна быть выше нуля")
            return

        if (angle_spin <= 0):
            messagebox.showerror("Ошибка", "Угол должен быть больше нуля")
            return

        p1 = [self.canvas_w // 2, self.canvas_h // 2]

        spin = 0

        while (spin <= 2 * pi):
            x2 = self.canvas_w // 2 + cos(spin) * line_len
            y2 = self.canvas_h // 2 + sin(spin) * line_len

            p2 = [x2, y2]

            self.parse_methods(p1, p2, self.methods_val.get(), self.color_code)

            spin += radians(angle_spin)

    def parse_line(self):
        try:
            x1 = int(self.x1_val.get())
            y1 = int(self.y1_val.get())
            x2 = int(self.x2_val.get())
            y2 = int(self.y2_val.get())
        except:
            messagebox.showerror("Ошибка", "Неверно введены координаты")
            return

        p1 = [x1, y1]
        p2 = [x2, y2]

        self.parse_methods(p1, p2, self.methods_val.get(), self.color_code)

    def parse_methods(self, p1, p2, option, option_color, draw=True):
        color = cu.Color(option_color)

        if (option == self.methods[0]):
            dots = self.cda_method(p1, p2, color)

            if draw:
                self.draw_line(dots)

        elif (option == self.methods[1]):
            dots = self.bresenham_float(p1, p2, color)

            if draw:
                self.draw_line(dots)

        elif (option == self.methods[2]):
            dots = self.bresenham_int(p1, p2, color)

            if draw:
                self.draw_line(dots)

        elif (option == self.methods[3]):
            dots = self.wu(p1, p2, color)

            if draw:
                self.draw_line(dots)

        elif (option == self.methods[4]):
            dots = self.bresenham_smooth(p1, p2, color)

            if draw:
                self.draw_line(dots)

        elif (option == self.methods[5]):
            self.lib_method(p1, p2, color)
        else:
            messagebox.showerror("Ошибка", "Неизвестный алгоритм")

    def wu(self, p1, p2, color, step_count=False):
        x1 = p1[0]
        y1 = p1[1]
        x2 = p2[0]
        y2 = p2[1]

        if (x2 - x1 == 0) and (y2 - y1 == 0):
            return [[x1, y1, color]]

        dx = x2 - x1
        dy = y2 - y1

        m = 1
        step = 1
        intens = 255

        dots = []

        steps = 0

        if (fabs(dy) > fabs(dx)):
            if (dy != 0):
                m = dx / dy
            m1 = m

            if (y1 > y2):
                m1 *= -1
                step *= -1

            y_end = round(y2) - 1 if (dy < dx) else (round(y2) + 1)

            for y_cur in range(round(y1), y_end, step):
                d1 = x1 - floor(x1)
                d2 = 1 - d1

                dot1 = [int(x1) + 1, y_cur, self.choose_color(color, round(fabs(d2) * intens))]

                dot2 = [int(x1), y_cur, self.choose_color(color, round(fabs(d1) * intens))]

                if step_count and y_cur < y2:
                    if (int(x1) != int(x1 + m)):
                        steps += 1

                dots.append(dot1)
                dots.append(dot2)

                x1 += m1

        else:
            if (dx != 0):
                m = dy / dx

            m1 = m

            if (x1 > x2):
                step *= -1
                m1 *= -1

            x_end = round(x2) - 1 if (dy > dx) else (round(x2) + 1)

            for x_cur in range(round(x1), x_end, step):
                d1 = y1 - floor(y1)
                d2 = 1 - d1

                dot1 = [x_cur, int(y1) + 1, self.choose_color(color, round(fabs(d2) * intens))]

                dot2 = [x_cur, int(y1), self.choose_color(color, round(fabs(d1) * intens))]

                if step_count and x_cur < x2:
                    if (int(y1) != int(y1 + m)):
                        steps += 1

                dots.append(dot1)
                dots.append(dot2)

                y1 += m1

        if step_count:
            return steps
        else:
            return dots

    def clear_canvas(self):
        self.canvas.delete("all")

    def lib_method(self, p1, p2, color):
        self.canvas.create_line(p1[0], p1[1], p2[0], p2[1], fill=color.hex)

    def cda_method(self, p1, p2, color, step_count=False):
        x1 = p1[0]
        y1 = p1[1]
        x2 = p2[0]
        y2 = p2[1]

        if (x2 - x1 == 0) and (y2 - y1 == 0):
            return [[x1, y1, color]]

        dx = x2 - x1
        dy = y2 - y1

        if (abs(dx) >= abs(dy)):
            l = abs(dx)
        else:
            l = abs(dy)

        dx /= l
        dy /= l

        x = round(x1)
        y = round(y1)

        dots = [[round(x), round(y), color]]

        i = 1

        steps = 0

        while (i < l):

            x += dx
            y += dy

            dot = [round(x), round(y), color]

            dots.append(dot)

            if step_count:
                if not ((round(x + dx) == round(x) and
                         round(y + dy) != round(y)) or
                        (round(x + dx) != round(x) and
                         round(y + dy) == round(y))):
                    steps += 1

            i += 1

        if step_count:
            return steps
        else:
            return dots

    def draw_line(self, dots):
        for dot in dots:
            self.canvas.create_line(dot[0], dot[1], dot[0] + 1, dot[1], fill=dot[2].hex)

    def sign(self, difference):
        if (difference < 0):
            return -1
        elif (difference == 0):
            return 0
        else:
            return 1

    def bresenham_float(self, p1, p2, color, step_count=False):
        x1 = p1[0]
        y1 = p1[1]
        x2 = p2[0]
        y2 = p2[1]

        if (x2 - x1 == 0) and (y2 - y1 == 0):
            return [[x1, y1, color]]

        x = x1
        y = y1

        dx = abs(x2 - x1)
        dy = abs(y2 - y1)

        s1 = self.sign(x2 - x1)
        s2 = self.sign(y2 - y1)

        if (dy > dx):
            tmp = dx
            dx = dy
            dy = tmp
            swaped = 1
        else:
            swaped = 0

        m = dy / dx
        e = m - 0.5
        i = 1

        dots = []

        steps = 0

        while (i <= dx + 1):
            dot = [x, y, color]
            dots.append(dot)

            x_buf = x
            y_buf = y

            while (e >= 0):
                if (swaped):
                    x = x + s1
                else:
                    y = y + s2

                e = e - 1

            if (swaped):
                y = y + s2
            else:
                x = x + s1

            e = e + m

            if step_count:
                if not ((x_buf == x and y_buf != y) or
                        (x_buf != x and y_buf == y)):
                    steps += 1

            i += 1

        if step_count:
            return steps
        else:
            return dots

    def bresenham_int(self, p1, p2, color, step_count=False):
        x1 = p1[0]
        y1 = p1[1]
        x2 = p2[0]
        y2 = p2[1]

        if (x2 - x1 == 0) and (y2 - y1 == 0):
            return [[x1, y1, color]]

        x = x1
        y = y1

        dx = abs(x2 - x1)
        dy = abs(y2 - y1)

        s1 = self.sign(x2 - x1)
        s2 = self.sign(y2 - y1)

        if (dy > dx):
            tmp = dx
            dx = dy
            dy = tmp
            swaped = 1
        else:
            swaped = 0

        e = 2 * dy - dx

        i = 1

        dots = []

        steps = 0

        while (i <= dx + 1):
            dot = [x, y, color]
            dots.append(dot)

            x_buf = x
            y_buf = y

            while (e >= 0):
                if (swaped):
                    x = x + s1
                else:
                    y = y + s2

                e = e - 2 * dx

            if (swaped):
                y = y + s2
            else:
                x = x + s1

            e = e + 2 * dy

            if step_count:
                if ((x_buf != x) and (y_buf != y)):
                    steps += 1

            i += 1

        if step_count:
            return steps
        else:
            return dots

    def choose_color(self, color, intens):
        return color + (intens, intens, intens)

    def bresenham_smooth(self, p1, p2, color, step_count=False):
        x1 = p1[0]
        y1 = p1[1]
        x2 = p2[0]
        y2 = p2[1]

        if (x2 - x1 == 0) and (y2 - y1 == 0):
            return [[x1, y1, color]]

        x = x1
        y = y1

        dx = abs(x2 - x1)
        dy = abs(y2 - y1)

        s1 = self.sign(x2 - x1)
        s2 = self.sign(y2 - y1)

        if (dy > dx):
            tmp = dx
            dx = dy
            dy = tmp
            swaped = 1
        else:
            swaped = 0

        intens = 255

        m = dy / dx
        e = intens / 2

        m *= intens
        w = intens - m

        dots = [[x, y, self.choose_color(color, round(e))]]

        i = 1

        steps = 0

        while (i <= dx):
            x_buf = x
            y_buf = y

            if (e < w):
                if (swaped):
                    y += s2
                else:
                    x += s1
                e += m
            else:
                x += s1
                y += s2

                e -= w

            dot = [x, y, self.choose_color(color, round(e))]

            dots.append(dot)

            if step_count:
                if not ((x_buf == x and y_buf != y) or
                        (x_buf != x and y_buf == y)):
                    steps += 1

            i += 1

        if step_count:
            return steps
        else:
            return dots

    def time_measure(self):
        time_mes = []

        try:
            line_len = float(self.len_val.get())
            angle_spin = float(self.ang_val.get())
        except:
            messagebox.showerror("Ошибка", "Неверно введены координаты")
            return

        if (line_len <= 0):
            messagebox.showerror("Ошибка", "Длина линии должна быть выше нуля")
            return

        if (angle_spin <= 0):
            messagebox.showerror("Ошибка", "Угол должен быть больше нуля")
            return

        for i in range(len(self.methods)):
            res_time = 0

            for _ in range(20):
                time_start = 0
                time_end = 0

                p1 = [self.canvas_w // 2, self.canvas_h // 2]

                spin = 0

                while (spin <= 2 * pi):
                    x2 = self.canvas_w // 2 + cos(spin) * line_len
                    y2 = self.canvas_h // 2 + sin(spin) * line_len

                    p2 = [x2, y2]

                    time_start += time.time()
                    self.parse_methods(p1, p2, self.methods[i], self.color_code, draw=False)
                    time_end += time.time()

                    spin += radians(angle_spin)

                res_time += (time_end - time_start)

                self.clear_canvas()

            time_mes.append(res_time / 20)

        plt.figure(figsize=(14, 6))

        plt.title("Замеры времени для различных методов")

        positions = np.arange(6)

        plt.xticks(positions, self.methods)
        plt.ylabel("Время")
        plt.bar(positions, time_mes, align="center", alpha=1)

        plt.show()

    def steps_measure(self):
        try:
            line_len = float(self.len_val.get())
        except:
            messagebox.showerror("Ошибка", "Неверно введены координаты")
            return

        if (line_len <= 0):
            messagebox.showerror("Ошибка", "Длина линии должна быть выше нуля")
            return

        p1 = [self.canvas_w // 2, self.canvas_h // 2]

        spin = 0

        angle_spin = [i for i in range(0, 91, 2)]

        cda_steps = []
        wu_steps = []
        bres_int_steps = []
        bres_float_steps = []
        bres_smooth_steps = []

        while (spin <= pi / 2 + 0.01):
            x2 = self.canvas_w // 2 + cos(spin) * line_len
            y2 = self.canvas_h // 2 + sin(spin) * line_len

            p2 = [x2, y2]

            cda_steps.append(self.cda_method(p1, p2, (255, 255, 255), step_count=True))
            wu_steps.append(self.wu(p1, p2, (255, 255, 255), step_count=True))
            bres_int_steps.append(self.bresenham_int(p1, p2, (255, 255, 255), step_count=True))
            bres_float_steps.append(self.bresenham_float(p1, p2, (255, 255, 255), step_count=True))
            bres_smooth_steps.append(self.bresenham_smooth(p1, p2, (255, 255, 255), step_count=True))

            spin += radians(2)

        plt.figure(figsize=(15, 6))

        plt.title("Замеры ступенчатости для различных методов\n{0} - длина отрезка".format(line_len))

        plt.xlabel("Угол (в градусах)")
        plt.ylabel("Количество ступенек")

        plt.plot(angle_spin, cda_steps, label="ЦДА")
        plt.plot(angle_spin, wu_steps, label="Ву")
        plt.plot(angle_spin, bres_float_steps, "-.", label="Брезенхем (float/int)")
        plt.plot(angle_spin, bres_smooth_steps, ":", label="Брезенхем\n(сглаживание)")

        plt.xticks(np.arange(91, step=5))

        plt.legend()

        plt.show()

    def exit(self):
        self.root.destroy()

    def __init__(self, root):
        self.root = root
        self.canvas_w = 1300
        self.canvas_h = 1000

        # Настройка параметров главного окна
        self.root.geometry(f"{int(root.winfo_screenwidth() * 0.76)}x{int(root.winfo_screenheight() * 0.85)}")
        self.root.minsize(int(root.winfo_screenwidth() * 0.76), int(root.winfo_screenheight() * 0.85))
        self.root.maxsize(int(root.winfo_screenwidth() * 0.76), int(root.winfo_screenheight() * 0.85))
        self.root.title("Geometry")

        # Создание левого фрейма
        self.up_frame = tk.Frame(root, background="grey")
        self.up_frame.pack(fill=tk.BOTH)

        main_point_label = tk.Label(self.up_frame, text="Выбор метода:", font=("system", 14))
        main_point_label.grid(padx=5, pady=5, column=0, row=0, rowspan=2)

        self.methods = ["ЦДА", "Брезенхем (float)", "Брезенхем (int)", "Ву", "Брезенхем (сглаживание)", "Библиотечная"]
        self.methods_val = tk.StringVar(value=self.methods[0])

        combobox = ttk.Combobox(self.up_frame, values=self.methods, textvariable=self.methods_val, font="system 14",
                                state="readonly", justify="center", width=30)
        combobox.grid(padx=5, pady=5, column=0, row=2, sticky=tk.EW)

        self.color_code = (0, 0, 0)
        self.bg_color_code = "white"
        color_button = tk.Button(self.up_frame, text="Выбрать\nцвет отрезка", command=self.set_color, font="system 14")
        color_button.grid(padx=5, pady=5, column=1, row=0, rowspan=2, sticky=tk.EW)
        bg_color_button = tk.Button(self.up_frame, text="Выбрать\nцвет фона", command=self.set_bg_color, font="system 14")
        bg_color_button.grid(padx=5, pady=5, column=1, row=2, sticky=tk.EW)

        x1_label = tk.Label(self.up_frame, text="x1:", font=("system", 14))
        x1_label.grid(padx=5, pady=5, column=2, row=0, sticky=tk.EW)

        self.x1_val = tk.StringVar()
        self.x1_val.set("")

        x1_entry = tk.Entry(self.up_frame, textvariable=self.x1_val, font=("system", 14), width=10)
        x1_entry.grid(padx=5, pady=5, column=3, row=0, sticky=tk.EW)

        y1_label = tk.Label(self.up_frame, text="y1:", font=("system", 14))
        y1_label.grid(padx=5, pady=5, column=4, row=0, sticky=tk.EW)

        self.y1_val = tk.StringVar()
        self.y1_val.set("")

        y1_entry = tk.Entry(self.up_frame, textvariable=self.y1_val, font=("system", 14), width=10)
        y1_entry.grid(padx=5, pady=5, column=5, row=0, sticky=tk.EW)

        x2_label = tk.Label(self.up_frame, text="x2:", font=("system", 14))
        x2_label.grid(padx=5, pady=5, column=2, row=1, sticky=tk.EW)

        self.x2_val = tk.StringVar()
        self.x2_val.set("")

        x2_entry = tk.Entry(self.up_frame, textvariable=self.x2_val, font=("system", 14), width=10)
        x2_entry.grid(padx=5, pady=5, column=3, row=1, sticky=tk.EW)

        y2_label = tk.Label(self.up_frame, text="y2:", font=("system", 14))
        y2_label.grid(padx=5, pady=5, column=4, row=1, sticky=tk.EW)

        self.y2_val = tk.StringVar()
        self.y2_val.set("")

        y2_entry = tk.Entry(self.up_frame, textvariable=self.y2_val, font=("system", 14), width=10)
        y2_entry.grid(padx=5, pady=5, column=5, row=1, sticky=tk.EW)

        draw_button = tk.Button(self.up_frame, text="Нарисовать отрезок", command=self.parse_line, font="system 14")
        draw_button.grid(padx=5, pady=5, column=2, row=2, columnspan=4, sticky=tk.EW)

        len_label = tk.Label(self.up_frame, text="Длина отрезка:", font=("system", 14))
        len_label.grid(padx=5, pady=5, column=6, row=0, sticky=tk.EW)

        self.len_val = tk.StringVar()
        self.len_val.set("")

        len_entry = tk.Entry(self.up_frame, textvariable=self.len_val, font=("system", 14), width=10)
        len_entry.grid(padx=5, pady=5, column=7, row=0, sticky=tk.EW)

        ang_label = tk.Label(self.up_frame, text="Угол:", font=("system", 14))
        ang_label.grid(padx=5, pady=5, column=6, row=1, sticky=tk.EW)

        self.ang_val = tk.StringVar()
        self.ang_val.set("")

        ang_entry = tk.Entry(self.up_frame, textvariable=self.ang_val, font=("system", 14), width=10)
        ang_entry.grid(padx=5, pady=5, column=7, row=1, sticky=tk.EW)

        draw_range_button = tk.Button(self.up_frame, text="Нарисовать спектр", command=self.parse_range, font="system 14")
        draw_range_button.grid(padx=5, pady=5, column=6, row=2, columnspan=2, sticky=tk.EW)

        time_button = tk.Button(self.up_frame, text="Сравнить время", command=self.time_measure, font="system 14")
        time_button.grid(padx=5, pady=5, column=8, row=0, sticky="nsew")

        steps_button = tk.Button(self.up_frame, text="Сравнить\nступенчатость", command=self.steps_measure, font="system 14")
        steps_button.grid(padx=5, pady=5, column=8, row=1, rowspan=2, sticky="nsew")

        clear_button = tk.Button(self.up_frame, text="Очистить\nэкран", command=self.clear_canvas, font="system 14")
        clear_button.grid(padx=5, pady=5, column=9, row=0, rowspan=2, sticky="nsew")

        exit_button = tk.Button(self.up_frame, text="Выход", command=self.exit, font="system 14")
        exit_button.grid(padx=5, pady=5, column=9, row=2, sticky="nsew")

        self.canvas = tk.Canvas(self.root, bg=self.bg_color_code, width = self.canvas_w, height = self.canvas_h)
        self.canvas.pack(padx=5, pady=5)
        
        len_entry.insert(tk.END, "400")
        ang_entry.insert(tk.END, "10")

        x1_entry.insert(tk.END, "0")
        y1_entry.insert(tk.END, "0")

        x2_entry.insert(tk.END, "500")
        y2_entry.insert(tk.END, "500")

if __name__ == "__main__":
    root = tk.Tk()
    app = GeometryApp(root)
    root.mainloop() 
    