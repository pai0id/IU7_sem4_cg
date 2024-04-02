import tkinter as tk  # Импорт модуля для создания графического интерфейса
import matplotlib.pyplot as plt
from matplotlib.backends.backend_tkagg import FigureCanvasTkAgg
from tkinter import messagebox  # Импорт модуля messagebox для вывода сообщений
from infowindow import InfoWindow  # Импорт InfoWindow из infowindow.py для вывода информационного окна
from calc import *  # Импорт всех функций из calc.py для расчетов
from matplotlib.patches import Polygon, Ellipse
import math  # Импорт модуля math для математических операций

class GeometryApp:
    def distance_between_points(self, p1, p2):
        return math.sqrt((p2[0] - p1[0]) ** 2 + (p2[1] - p1[1]) ** 2)

    def parabola_equation(self, x):
        return self.c - (0.*x - self.d)**2

    def upper_circle_equation(self, x):
        delt = self.r**2 - (x - self.a)**2
        return np.sqrt(delt) + self.b

    def lower_circle_equation(self, x):
        delt = self.r ** 2 - (x - self.a) ** 2
        return -np.sqrt(delt) + self.b

    def get_contour(self):
        circle_points = []
        parabola_points = []

        y_parabola = self.parabola_equation(self.circle_x)
        y_lower_cir = self.lower_circle_equation(self.circle_x)
        y_upper_cir = self.upper_circle_equation(self.circle_x)

        for i in range(len(self.circle_x)):
            if y_lower_cir[i] <= y_parabola[i] <= y_upper_cir[i]:
                parabola_points.append((self.circle_x[i], y_parabola[i]))
            if y_lower_cir[i] <= y_parabola[i]:
                circle_points.append((self.circle_x[i], y_lower_cir[i]))
            if y_upper_cir[i] <= y_parabola[i]:
                circle_points.append((self.circle_x[i], y_upper_cir[i]))

        return parabola_points + circle_points[::-1]

    def check_mid(self, before, after, val):
        return min(before, after) < val < max(before, after)

    def plot_graph(self, ax):
        ax.plot(self.circle[0][1:], self.circle[1][1:], 'g-')
        plt.scatter(self.circle[0][0], self.circle[1][0], color="black")
        ax.text(self.circle[0][0], self.circle[1][0] + 0.1, f"({self.circle[0][0]:.2f}, {self.circle[1][0]:.2f})")

        ax.plot(self.parabola[0], self.parabola[1], 'g-')

        if self.intersection_points:
            intersection_polygon = Polygon(np.array(self.intersection_points), color='green', alpha=0.9)
            ax.add_patch(intersection_polygon)

    def validate_input(self, new_value):
        if new_value == "" or new_value == "-":
            return True
        try:
            float(new_value)
            return True
        except ValueError:
            return False

    def on_validate(self, p):
        if self.validate_input(p):
            return True
        else:
            return False

    def banana_rotate(self):
        try:
            x = float(self.main_point_x_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите x точки поворота")
            return
        try:
            y = float(self.main_point_y_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите y точки поворота")
            return
        try:
            angle = float(self.rotate_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите угол поворота")
            return
        self.circle = plot_rotate(self.circle, (x, y), angle)
        self.parabola = plot_rotate(self.parabola, (x, y), angle)
        self.intersection_points = alt_plot_rotate(self.intersection_points, (x, y), angle)
        self.log.append((self.circle, self.parabola, self.intersection_points))
        self.reset_canvas(self.ax)
        plt.scatter(x, y, color="black")
        self.ax.text(x, y + 0.1, f"({x:.2f}, {y:.2f})")
        self.plot_graph(self.ax)
        self.canvas.draw()

    def banana_scale(self):
        try:
            x = float(self.main_point_x_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите x точки масштабирования")
            return
        try:
            y = float(self.main_point_y_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите y точки масштабирования")
            return
        try:
            kx = float(self.scale_x_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите kx масштабирования")
            return
        try:
            ky = float(self.scale_y_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите ky масштабирования")
            return
        self.circle = plot_scale(self.circle, (x, y), kx, ky)
        self.parabola = plot_scale(self.parabola, (x, y), kx, ky)
        self.intersection_points = alt_plot_scale(self.intersection_points, (x, y), kx, ky)
        self.log.append((self.circle, self.parabola, self.intersection_points))
        self.reset_canvas(self.ax)
        plt.scatter(x, y, color="black")
        self.ax.text(x, y + 0.1, f"({x:.2f}, {y:.2f})")
        self.plot_graph(self.ax)
        self.canvas.draw()

    def banana_move(self):
        try:
            dx = float(self.move_x_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите dx")
            return
        try:
            dy = float(self.move_y_val.get())
        except Exception:
            messagebox.showerror("Ошибка", f"Введите dy")
            return
        self.circle = plot_move(self.circle, dx, dy)
        self.parabola = plot_move(self.parabola, dx, dy)
        self.intersection_points = alt_plot_move(self.intersection_points, dx, dy)
        self.log.append((self.circle, self.parabola, self.intersection_points))
        self.reset_canvas(self.ax)
        self.plot_graph(self.ax)
        self.canvas.draw()

    def change_param(self):
        ask_window = tk.Toplevel(self.root)
        ask_window.title("Параметры")
        ask_window.geometry("655x200+1100+200")

        a_label = tk.Label(ask_window, text="A:", font=("system", 14))
        a_label.grid(padx=5, pady=5, column=0, row=0)

        self.a_val.set("")

        a_entry = tk.Entry(ask_window, textvariable=self.a_val, font=("system", 14),
                                      validate="key", validatecommand=(self.vcmd, '%P'))
        a_entry.grid(padx=5, pady=5, column=1, row=0,)

        b_label = tk.Label(ask_window, text="B:", font=("system", 14))
        b_label.grid(padx=5, pady=5, column=2, row=0)

        self.b_val.set("")

        b_entry = tk.Entry(ask_window, textvariable=self.b_val, font=("system", 14),
                           validate="key", validatecommand=(self.vcmd, '%P'))
        b_entry.grid(padx=5, pady=5, column=3, row=0)

        r_label = tk.Label(ask_window, text="R:", font=("system", 14))
        r_label.grid(padx=5, pady=5, column=0, columnspan=2, row=1)

        self.r_val.set("")

        r_entry = tk.Entry(ask_window, textvariable=self.r_val, font=("system", 14),
                           validate="key", validatecommand=(self.vcmd, '%P'))
        r_entry.grid(padx=5, pady=5, column=2, columnspan=2, row=1)

        c_label = tk.Label(ask_window, text="C:", font=("system", 14))
        c_label.grid(padx=5, pady=5, column=0, row=2)

        self.c_val.set("")

        c_entry = tk.Entry(ask_window, textvariable=self.c_val, font=("system", 14),
                           validate="key", validatecommand=(self.vcmd, '%P'))
        c_entry.grid(padx=5, pady=5, column=1, row=2)

        d_label = tk.Label(ask_window, text="D:", font=("system", 14))
        d_label.grid(padx=5, pady=5, column=2, row=2)

        self.d_val.set("")

        d_entry = tk.Entry(ask_window, textvariable=self.d_val, font=("system", 14),
                           validate="key", validatecommand=(self.vcmd, '%P'))
        d_entry.grid(padx=5, pady=5, column=3, row=2)

        apply_changes = tk.Button(ask_window, text="Применить изменения", command=self.apply_changes, font=("system", 14))
        apply_changes.grid(padx=5, pady=5, column=3, row=3)

    def apply_changes(self):
        try:
            a = float(self.a_val.get())
            b = float(self.b_val.get())
            r = float(self.r_val.get())
            c = float(self.c_val.get())
            d = float(self.d_val.get())
        except Exception:
            messagebox.showerror("Ошибка", "Введите все параметры")
            return
        if r <= 0:
            messagebox.showerror("Ошибка", "Радиус должен быть больше 0")
            return
        self.a = a
        self.b = b
        self.r = r
        self.c = c
        self.d = d

        self.x_range = (self.a - 4 * self.r, self.a + 4 * self.r) if self.r != 0 else (self.a - 1, self.a + 1)
        self.x = np.linspace(self.x_range[0], self.x_range[1], 1000)
        self.y_range = (self.b - 4 * self.r, self.b + 4 * self.r) if self.r != 0 else (self.b - 1, self.b + 1)

        self.circle_x = np.linspace(self.a - self.r, self.a + self.r, 1000)
        self.circle = [np.concatenate([np.array([self.a]), self.circle_x, self.circle_x[::-1]]),
                       np.concatenate([np.array([self.b]), self.upper_circle_equation(self.circle_x),
                                       self.lower_circle_equation(self.circle_x)[::-1]])]
        self.parabola = [self.x, self.parabola_equation(self.x)]
        self.intersection_points = self.get_contour()

        self.log = [(self.circle, self.parabola, self.intersection_points)]

        self.reset_canvas(self.ax)
        self.plot_graph(self.ax)
        self.reset_canvas(self.const_ax)
        self.plot_graph(self.const_ax)
        self.const_canvas.draw()
        self.canvas.draw()

    def info(self):
        # Отображение информации о программе в новом окне
        InfoWindow(self.root, "Приложение рисует графики (x-a)^2+(y-b)^2=r^2 и y=c-(x-d)^2 в декартовой системе координат, закрашивает их пересечение и выполняет преобразования над ними."
                              "\nНачальные параметры:\na=0 b=0 r=5\nc=0 d=0"
                              "\n Поляков Андрей ИУ7-42Б")

    def undo(self):
        if len(self.log) <= 1:
            messagebox.showerror("Ошибка", "Исходное изображение достигнуто")
            return
        self.circle, self.parabola, self.intersection_points = self.log[-2]
        self.log.pop()
        self.reset_canvas(self.ax)
        self.plot_graph(self.ax)
        self.canvas.draw()

    def reset(self):
        if len(self.log) <= 1:
            return
        self.circle, self.parabola, self.intersection_points = self.log[0]
        self.log.clear()
        self.log.append((self.circle, self.parabola, self.intersection_points))
        self.reset_canvas(self.ax)
        self.plot_graph(self.ax)
        self.canvas.draw()

    def exit(self):
        # Закрытие приложения
        self.root.destroy()

    def reset_canvas(self, ax):
        ax.clear()
        ax.set_xlim(self.x_range)
        ax.set_ylim(self.y_range)
        ax.set_xlabel('X')
        ax.set_ylabel('Y')
        ax.grid(alpha=0.4)

    def __init__(self, root):
        self.root = root
        self.vcmd = self.root.register(self.on_validate)
        self.a = 0
        self.a_val = tk.StringVar()
        self.b = 0
        self.b_val = tk.StringVar()
        self.r = 5
        self.r_val = tk.StringVar()
        self.c = 0
        self.c_val = tk.StringVar()
        self.d = 0
        self.d_val = tk.StringVar()
        self.x_range = (self.a - 4 * self.r, self.a + 4 * self.r) if self.r != 0 else (self.a - 1, self.a + 1)
        self.x = np.linspace(self.x_range[0], self.x_range[1], 1000)
        self.y_range = (self.b - 4 * self.r, self.b + 4 * self.r) if self.r != 0 else (self.b - 1, self.b + 1)

        self.circle_x = np.linspace(self.a - self.r, self.a + self.r, 1000)
        self.circle = [np.concatenate([np.array([self.a]), self.circle_x, self.circle_x[::-1]]),
                       np.concatenate([np.array([self.b]), self.upper_circle_equation(self.circle_x),
                                                self.lower_circle_equation(self.circle_x)[::-1]])]
        self.parabola = [self.x, self.parabola_equation(self.x)]
        self.intersection_points = self.get_contour()

        self.log = [(self.circle, self.parabola, self.intersection_points)]

        # Настройка параметров главного окна
        self.root.geometry(f"{int(root.winfo_screenwidth() * 0.9)}x{int(root.winfo_screenheight() * 0.9)}")
        self.root.minsize(int(root.winfo_screenwidth() * 0.9), int(root.winfo_screenheight() * 0.9))
        self.root.maxsize(int(root.winfo_screenwidth() * 0.9), int(root.winfo_screenheight() * 0.9))
        self.root.title("Geometry")

        # Создание левого фрейма
        self.left_frame = tk.Frame(root, background="grey")
        self.left_frame.pack(side=tk.LEFT, fill=tk.BOTH)

        main_point_label = tk.Label(self.left_frame, text="Точка масштабирования/поворота", font=("system", 18),
                                    background="grey", width=40)
        main_point_label.grid(padx=5, pady=5, column=0, row=0, columnspan=4, sticky=tk.EW)

        main_point_x_label = tk.Label(self.left_frame, text="X:", font=("system", 14))
        main_point_x_label.grid(padx=5, pady=5, column=0, row=1, sticky=tk.EW)

        self.main_point_x_val = tk.StringVar()
        self.main_point_x_val.set("")

        main_point_x_entry = tk.Entry(self.left_frame, textvariable=self.main_point_x_val, font=("system", 14),
                                      validate="key", validatecommand=(self.vcmd, '%P'))
        main_point_x_entry.grid(padx=5, pady=5, column=1, row=1, sticky=tk.EW)

        main_point_y_label = tk.Label(self.left_frame, text="Y:", font=("system", 14))
        main_point_y_label.grid(padx=5, pady=5, column=2, row=1, sticky=tk.EW)

        self.main_point_y_val = tk.StringVar()
        self.main_point_y_val.set("")

        main_point_y_entry = tk.Entry(self.left_frame, textvariable=self.main_point_y_val, font=("system", 14),
                                      validate="key", validatecommand=(self.vcmd, '%P'))
        main_point_y_entry.grid(padx=5, pady=5, column=3, row=1, sticky=tk.EW)

        move_label = tk.Label(self.left_frame, text="Перемещение", background="grey", font=("system", 14))
        move_label.grid(padx=5, pady=5, column=0, row=2, columnspan=2, sticky=tk.EW)

        move_button = tk.Button(self.left_frame, text="Переместить", command=self.banana_move, font=("system", 14))
        move_button.grid(padx=5, pady=5, column=2, row=2, columnspan=2, sticky=tk.EW)

        move_x_label = tk.Label(self.left_frame, text="dx:", font=("system", 14))
        move_x_label.grid(padx=5, pady=5, column=0, row=3, sticky=tk.EW)

        self.move_x_val = tk.StringVar()
        self.move_x_val.set("")

        move_x_entry = tk.Entry(self.left_frame, textvariable=self.move_x_val, font=("system", 14),
                                      validate="key", validatecommand=(self.vcmd, '%P'))
        move_x_entry.grid(padx=5, pady=5, column=1, row=3, sticky=tk.EW)

        move_y_label = tk.Label(self.left_frame, text="dy:", font=("system", 14))
        move_y_label.grid(padx=5, pady=5, column=2, row=3, sticky=tk.EW)

        self.move_y_val = tk.StringVar()
        self.move_y_val.set("")

        move_y_entry = tk.Entry(self.left_frame, textvariable=self.move_y_val, font=("system", 14),
                                      validate="key", validatecommand=(self.vcmd, '%P'))
        move_y_entry.grid(padx=5, pady=5, column=3, row=3, sticky=tk.EW)

        scale_label = tk.Label(self.left_frame, text="Масштабирование", background="grey", font=("system", 14))
        scale_label.grid(padx=5, pady=5, column=0, row=4, columnspan=2, sticky=tk.EW)

        scale_button = tk.Button(self.left_frame, text="Масштабировать", command=self.banana_scale, font=("system", 14))
        scale_button.grid(padx=5, pady=5, column=2, row=4, columnspan=2, sticky=tk.EW)

        scale_x_label = tk.Label(self.left_frame, text="kx:", font=("system", 14))
        scale_x_label.grid(padx=5, pady=5, column=0, row=5, sticky=tk.EW)

        self.scale_x_val = tk.StringVar()
        self.scale_x_val.set("")

        scale_x_entry = tk.Entry(self.left_frame, textvariable=self.scale_x_val, font=("system", 14),
                                 validate="key", validatecommand=(self.vcmd, '%P'))
        scale_x_entry.grid(padx=5, pady=5, column=1, row=5, sticky=tk.EW)

        scale_y_label = tk.Label(self.left_frame, text="ky:", font=("system", 14))
        scale_y_label.grid(padx=5, pady=5, column=2, row=5, sticky=tk.EW)

        self.scale_y_val = tk.StringVar()
        self.scale_y_val.set("")

        scale_y_entry = tk.Entry(self.left_frame, textvariable=self.scale_y_val, font=("system", 14),
                                 validate="key", validatecommand=(self.vcmd, '%P'))
        scale_y_entry.grid(padx=5, pady=5, column=3, row=5, sticky=tk.EW)

        rotate_label = tk.Label(self.left_frame, text="Поворот", background="grey", font=("system", 14))
        rotate_label.grid(padx=5, pady=5, column=0, row=6, columnspan=2, sticky=tk.EW)

        rotate_button = tk.Button(self.left_frame, text="Повернуть", command=self.banana_rotate, font=("system", 14))
        rotate_button.grid(padx=5, pady=5, column=2, row=6, columnspan=2, sticky=tk.EW)

        rotate_label = tk.Label(self.left_frame, text="Угол:", font=("system", 14))
        rotate_label.grid(padx=5, pady=5, column=0, columnspan=2, row=7, sticky=tk.EW)

        self.rotate_val = tk.StringVar()
        self.rotate_val.set("")

        rotate_entry = tk.Entry(self.left_frame, textvariable=self.rotate_val, font=("system", 14),
                                validate="key", validatecommand=(self.vcmd, '%P'))
        rotate_entry.grid(padx=5, pady=5, column=2, columnspan=2, row=7, sticky=tk.EW)

        change_param_button = tk.Button(self.left_frame, text="Изменить параметры уравнений",
                                        command=self.change_param, font=("system", 14))
        change_param_button.grid(padx=5, pady=5, column=0, row=8, columnspan=4, sticky=tk.EW)

        info_button = tk.Button(self.left_frame, text="Информация",
                                        command=self.info, font=("system", 14))
        info_button.grid(padx=5, pady=5, column=0, row=9, columnspan=4, sticky=tk.EW)

        undo_button = tk.Button(self.left_frame, text="Шаг назад",
                                command=self.undo, font=("system", 14))
        undo_button.grid(padx=5, pady=5, column=0, row=10, columnspan=4, sticky=tk.EW)

        reset_button = tk.Button(self.left_frame, text="Исходное изображение",
                                command=self.reset, font=("system", 14))
        reset_button.grid(padx=5, pady=5, column=0, row=11, columnspan=4, sticky=tk.EW)

        exit_button = tk.Button(self.left_frame, text="Выход",
                                 command=self.exit, font=("system", 14))
        exit_button.grid(padx=5, pady=5, column=0, row=12, columnspan=4, sticky=tk.EW)

        # Создание нижнего фрейма
        self.low_frame = tk.Frame(self.left_frame, background="grey")
        self.low_frame.grid(column=0, row=13, columnspan=4, sticky=tk.EW)

        const_fig, self.const_ax = plt.subplots()
        self.const_canvas = FigureCanvasTkAgg(const_fig, master=self.low_frame)
        self.reset_canvas(self.const_ax)
        self.plot_graph(self.const_ax)
        self.const_canvas.draw()
        self.const_canvas.get_tk_widget().pack(expand=1)

        tk.Label(self.left_frame, text="Исходное изображение", background="grey", font=("system", 14)).grid(column=0, row=14, columnspan=4, sticky=tk.EW)

        # Создание правого фрейма
        self.right_frame = tk.Frame(root, background="grey")
        self.right_frame.pack(side=tk.RIGHT, fill=tk.BOTH, expand=1)

        fig, self.ax = plt.subplots()
        self.canvas = FigureCanvasTkAgg(fig, master=self.right_frame)
        self.reset_canvas(self.ax)
        self.plot_graph(self.ax)
        self.canvas.draw()
        self.canvas.get_tk_widget().pack(expand=1, fill=tk.BOTH)
