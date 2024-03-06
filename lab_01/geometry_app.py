import tkinter as tk  # Импорт модуля для создания графического интерфейса
from tkinter import ttk  # Импорт модуля ttk для дополнительных графических элементов
from tkinter import messagebox  # Импорт модуля messagebox для вывода сообщений
from tkinter.scrolledtext import ScrolledText  # Импорт модуля ScrolledText для создания прокручиваемого текстового поля
from infowindow import InfoWindow  # Импорт InfoWindow из infowindow.py для вывода информационного окна
from calc import *  # Импорт всех функций из calc.py для расчетов
import math  # Импорт модуля math для математических операций

class GeometryApp:
    def __init__(self, root):
        self.root = root
        self.points = []  # Список для хранения точек

        # Настройка параметров главного окна
        self.root.geometry(f"{int(root.winfo_screenwidth() * 0.85)}x{int(root.winfo_screenheight() * 0.79) - 8}")
        self.root.minsize(int(root.winfo_screenwidth() * 0.85), int(root.winfo_screenheight() * 0.79) - 8)
        self.root.maxsize(int(root.winfo_screenwidth() * 0.85), int(root.winfo_screenheight() * 0.79) - 8)
        self.root.title("Geometry")

        # Создание левого фрейма
        self.left_frame = tk.Frame(root, width=int(root.winfo_screenwidth() / 3), background="grey")
        self.left_frame.pack(side=tk.LEFT, fill=tk.Y)

        # Создание элементов для ввода координат точек
        self.x_label = tk.Label(self.left_frame, text="X:", font="system")
        self.x_label.grid(padx=5, pady=5, column=0, row=0, sticky=tk.EW)

        self.x_val = tk.StringVar()
        self.x_val.set("")

        self.x_entry = tk.Entry(self.left_frame, textvariable=self.x_val, font="system")
        self.x_entry.grid(padx=5, pady=5, column=1, row=0, sticky=tk.EW)

        self.y_label = tk.Label(self.left_frame, text="Y:", font="system")
        self.y_label.grid(padx=5, pady=5, column=2, row=0, sticky=tk.EW)

        self.y_val = tk.StringVar()
        self.y_val.set("")

        self.y_entry = tk.Entry(self.left_frame, textvariable=self.y_val, font="system")
        self.y_entry.grid(padx=5, pady=5, column=3, row=0, sticky=tk.EW)

        # Кнопка для добавления точки
        self.add_button = tk.Button(self.left_frame, text="Добавить", command=self.add_point, width=25, font="system")
        self.add_button.grid(padx=5, pady=5, column=0, columnspan=2, row=1, sticky=tk.EW)

        # Кнопка для удаления точки
        self.del_button = tk.Button(self.left_frame, text="Удалить", command=self.del_point, width=25, font="system")
        self.del_button.grid(padx=5, pady=5, column=2, columnspan=2, row=1, sticky=tk.EW)

        # Создание таблицы для отображения списка точек
        self.table_columns = ("num", "x", "y")
        self.table = ttk.Treeview(self.left_frame, columns=self.table_columns, show="headings")
        self.table.grid(padx=5, pady=5, column=0, row=2, columnspan=4, sticky=tk.NSEW)
        # Настройка заголовков столбцов
        self.table.heading("num", text="Номер")
        self.table.heading("x", text="x")
        self.table.heading("y", text="y")
        # Установка якорей для столбцов
        self.table.column("#1", anchor=tk.N)
        self.table.column("#2", anchor=tk.N)
        self.table.column("#3", anchor=tk.N)

        # Создание вертикального скроллбара для таблицы
        self.scrollbar = ttk.Scrollbar(self.left_frame, orient=tk.VERTICAL, command=self.table.yview)
        self.table.configure(yscroll=self.scrollbar.set)
        self.scrollbar.grid(row=2, column=3, sticky="nes")

        # Кнопка для удаления всех точек
        self.del_all_button = tk.Button(self.left_frame, text="Удалить все точки", command=self.del_all, width=60, font="system")
        self.del_all_button.grid(padx=5, pady=5, column=0, columnspan=4, row=3, sticky=tk.NSEW)

        # Кнопка для редактирования точек
        self.edit_button = tk.Button(self.left_frame, text="Редактировать точки", command=self.edit, font="system")
        self.edit_button.grid(padx=5, pady=5, column=0, columnspan=2, sticky=tk.EW, row=4)

        # Кнопка для вывода информации о задаче
        self.info_button = tk.Button(self.left_frame, text="Вывести информацию о задаче", command=self.info, font="system")
        self.info_button.grid(padx=5, pady=5, column=2, columnspan=2, sticky=tk.EW, row=4)

        # Кнопка для вывода результатов расчета
        self.res_button = tk.Button(self.left_frame, text="Вывести результат", command=self.calc, font="system", background="#FFC33B", activebackground="#FFD168")
        self.res_button.grid(padx=5, pady=5, column=0, sticky=tk.EW, columnspan=4, row=5)

        # Создание надписи "Журнал"
        self.log_label = tk.Label(self.left_frame, text="Журнал", font="system", background="white")
        self.log_label.grid(padx=5, pady=5, column=0, columnspan=4, row=6)

        # Кнопка для выхода из приложения
        self.exit_button = tk.Button(self.left_frame, text="Выход", command=self.exit)
        self.exit_button.grid(padx=5, pady=5, column=0, columnspan=2, row=6, sticky=tk.W)

        # Создание прокручиваемого текстового поля для журнала
        self.log = ScrolledText(self.left_frame, font="system", height=19, width=60, bg="white", state=tk.DISABLED)
        self.log.grid(padx=5, pady=5, column=0, columnspan=4, row=7, sticky=tk.NSEW)

        # Создание правого фрейма
        self.right_frame = tk.Frame(root, background="grey")
        self.right_frame.pack(side=tk.RIGHT, fill=tk.Y)

        # Создание холста для рисования геометрических фигур
        self.canvas = tk.Canvas(self.right_frame, width=int(root.winfo_screenwidth() * 2 / 3), height=int(root.winfo_screenheight()), bg="white")
        self.canvas.pack(padx=10, pady=10)

    def get_pair(self, id):
        # Получить пару координат для указанного идентификатора точки
        return [float(self.table.set(id, 1)), float(self.table.set(id, 2))]

    def add_point(self):
        try:
            # Попытка получить значения X и Y из полей ввода
            x = float(self.x_val.get())
            y = float(self.y_val.get())
        except Exception as e:
            # Обработка исключения при некорректном формате ввода
            messagebox.showerror("Ошибка", f"Некорректный формат")
            return
        # Создание точки и добавление ее в таблицу
        dot = (len(self.points) + 1, x, y)
        self.points.append([x, y])
        self.table.insert("", tk.END, values=dot)
        self.table.yview_moveto(1)
        self.log_add(f"Добавлена точка ({x}, {y})\n")

    def del_point(self):
        # Удаление выделенных точек из таблицы и очистка списка точек
        for item in self.table.selection():
            self.log_add(f"Удалена точка ({self.table.set(item, 1)}, {self.table.set(item, 2)})\n")
            self.table.delete(item)
        self.points.clear()
        children = self.table.get_children()
        # Обновление списка точек после удаления
        for i in range(len(children)):
            self.table.set(children[i], 0, i + 1)
            self.points.append(self.get_pair(children[i]))
            
        self.calc()

    def del_all(self):
        # Удаление всех точек из таблицы и списка точек
        for item in self.table.get_children():
            self.log_add(f"Удалена точка ({self.table.set(item, 1)}, {self.table.set(item, 2)})\n")
            self.table.delete(item)
        self.points.clear()
        
        self.canvas.delete(tk.ALL)

    def info(self):
        # Отображение информации о программе в новом окне
        InfoWindow(self.root, "    Программа ищет среди множества точек на плоскости такие три, \
которые образуют треугольник, для которого разность площадей описанного и вписанного кругов максимальна.\
\n    Программа позволяет добавлять точки, вводя их координаты в поля 'X' и 'Y', удалять точки, выделив их в таблице, \
а также редактировать точки, заменяя координаты всех выделенных точек на введенные.")

    def draw_triangle(self, triangle):
        # Отображение треугольника на холсте
        self.canvas.create_polygon(triangle, fill='', outline='black', width=3)

    def draw_circle(self, center, radius, fill, outline, stipple=""):
        # Отображение круга на холсте
        x, y = center
        self.canvas.create_oval(x - radius, y - radius, x + radius, y + radius, outline=outline, fill=fill, width=3,
                                stipple=stipple)

    def find_left_x(self):
        # Нахождение наименьшей координаты X
        x_arr = [self.triangle[0][0], self.triangle[1][0], self.triangle[2][0],
                 self.inscribed_circle[0][0] - self.inscribed_circle[1], self.circumscribed_circle[0][0] - self.circumscribed_circle[1]]
        return min(x_arr)

    def find_right_x(self):
        # Нахождение наибольшей координаты X
        x_arr = [self.triangle[0][0], self.triangle[1][0], self.triangle[2][0],
                 self.inscribed_circle[0][0] + self.inscribed_circle[1], self.circumscribed_circle[0][0] + self.circumscribed_circle[1]]
        return max(x_arr)

    def find_up_y(self):
        # Нахождение наибольшей координаты Y
        y_arr = [self.triangle[0][1], self.triangle[1][1], self.triangle[2][1],
                 self.inscribed_circle[0][1] + self.inscribed_circle[1], self.circumscribed_circle[0][1] + self.circumscribed_circle[1]]
        return max(y_arr)

    def find_down_y(self):
        # Нахождение наименьшей координаты Y
        y_arr = [self.triangle[0][1], self.triangle[1][1], self.triangle[2][1],
                 self.inscribed_circle[0][1] - self.inscribed_circle[1], self.circumscribed_circle[0][1] - self.circumscribed_circle[1]]
        return min(y_arr)

    def translate_x_to_canvas(self, val):
        # Перевод координаты X в координаты холста
        return int((val - self.center[0]) * self.scale) + 50

    def translate_y_to_canvas(self, val):
        # Перевод координаты Y в координаты холста
        return int((self.center[1] - val) * self.scale) + 50

    def translate_circle_to_canvas(self, circle):
        # Перевод координат и радиуса круга в координаты холста
        return ((self.translate_x_to_canvas(circle[0][0]), self.translate_y_to_canvas(circle[0][1])),
                circle[1] * self.scale)

    def translate_poly_to_canvas(self, poly):
        # Перевод координат вершин полигона в координаты холста
        dots = []
        for dot in poly:
            dots.append([self.translate_x_to_canvas(dot[0]), self.translate_y_to_canvas(dot[1])])
        return dots

    def calc(self):
        # Вычисление и отображение треугольника, вписанного и описанного кругов
        self.triangle, c_area, i_area = find_triangle(self.points)
        if not self.triangle:
            # Проверка на наличие треугольника
            messagebox.showerror("Ошибка", f"Треугольник не найден")
            return

        dots = [0, 0, 0]

        for item in self.table.get_children():
            if dots[0] and dots[1] and dots[2]:
                break
            if not dots[0] and float_eq(float(self.table.set(item, 1)), self.triangle[0][0]) and float_eq(
                    float(self.table.set(item, 2)), self.triangle[0][1]):
                dots[0] = self.table.set(item, 0)
            if not dots[1] and float_eq(float(self.table.set(item, 1)), self.triangle[1][0]) and float_eq(
                    float(self.table.set(item, 2)), self.triangle[1][1]):
                dots[1] = self.table.set(item, 0)
            if not dots[2] and float_eq(float(self.table.set(item, 1)), self.triangle[2][0]) and float_eq(
                    float(self.table.set(item, 2)), self.triangle[2][1]):
                dots[2] = self.table.set(item, 0)

        # Вывод информации о треугольнике и площадях кругов
        self.log_add(f"\n-----------------------------------\n")
        self.log_add(f"Искомый треугольник образуют точки:\n")
        self.log_add(f"{dots[0]} - ({self.triangle[0][0]}, {self.triangle[0][1]})\n")
        self.log_add(f"{dots[1]} - ({self.triangle[1][0]}, {self.triangle[1][1]})\n")
        self.log_add(f"{dots[2]} - ({self.triangle[2][0]}, {self.triangle[2][1]})\n")
        self.log_add(f"Площадь вписанного круга: {i_area:.6f}\n")
        self.log_add(f"Площадь описанного круга: {c_area:.6f}\n")
        self.log_add(f"Разность площадей: {(c_area - i_area):.6f}\n")
        self.log_add(f"-----------------------------------\n\n")

        self.inscribed_circle = find_incircle(self.triangle)
        self.circumscribed_circle = find_circumcircle(self.triangle)

        # Нахождение границ рисунка и масштабирование
        x_min = self.find_left_x()
        x_max = self.find_right_x()
        y_min = self.find_down_y()
        y_max = self.find_up_y()
        x_delt = x_max - x_min
        y_delt = y_max - y_min
        self.center = (x_min, y_max)

        self.root.update()

        if x_delt < y_delt:
            self.scale = (self.canvas.winfo_width() - 200) / x_delt
        else:
            self.scale = (self.canvas.winfo_height() - 200) / y_delt

        self.canvas.delete(tk.ALL)

        # Отображение треугольника и кругов на холсте
        self.draw_triangle(self.translate_poly_to_canvas(self.triangle))
        new_circle = self.translate_circle_to_canvas(self.circumscribed_circle)
        self.draw_circle(new_circle[0], new_circle[1], "#B62000", "#B62000", "gray25")
        new_circle = self.translate_circle_to_canvas(self.inscribed_circle)
        self.draw_circle(new_circle[0], new_circle[1], "#02AA00", "#02AA00", "gray25")
        
        centers = [self.inscribed_circle[0], self.circumscribed_circle[0]]
        rads = [self.inscribed_circle[1], self.circumscribed_circle[1]]
        for i in range(len(centers)):
            self.draw_circle((self.translate_x_to_canvas(centers[i][0]),
                              self.translate_y_to_canvas(centers[i][1])), 5, "black", "black")
            txt = self.canvas.create_text(self.translate_x_to_canvas(centers[i][0]) + 20,
                                          self.translate_y_to_canvas(centers[i][1]) + 30,
                                          text=f"({centers[i][0]:.2f}, {centers[i][1]:.2f})\nR = {rads[i]:.2f}")
            rct = self.canvas.create_rectangle(self.canvas.bbox(txt), fill="white")
            self.canvas.tag_lower(rct, txt)

        for i in range(len(self.triangle)):
            self.draw_circle((self.translate_x_to_canvas(self.triangle[i][0]),
                              self.translate_y_to_canvas(self.triangle[i][1])), 5, "black", "black")
            txt = self.canvas.create_text(self.translate_x_to_canvas(self.triangle[i][0]) + 20,
                                          self.translate_y_to_canvas(self.triangle[i][1]) + 20,
                                          text=f"{dots[i]} - ({self.triangle[i][0]:.2f}, {self.triangle[i][1]:.2f})")
            rct = self.canvas.create_rectangle(self.canvas.bbox(txt), fill="white")
            self.canvas.tag_lower(rct, txt)

    def edit(self):
        try:
            # Попытка получить новые значения координат точек
            x = float(self.x_val.get())
            y = float(self.y_val.get())
        except Exception as e:
            # Обработка исключения при некорректном формате ввода
            messagebox.showerror("Ошибка", f"Некорректный формат")
            return
        # Редактирование координат точек
        for item in self.table.selection():
            self.log_add(f"Точка {self.table.set(item, 0)} теперь: ({x}, {y})\n")
            self.table.set(item, 1, x)
            self.table.set(item, 2, y)
            self.points[int(self.table.set(item, 0)) - 1] = [x, y]
        self.calc()

    def log_add(self, string):
        # Добавление сообщения в лог
        self.log.configure(state=tk.NORMAL)
        self.log.insert(tk.END, string)
        self.log.configure(state=tk.DISABLED)
        self.log.see(tk.END)

    def exit(self):
        # Закрытие приложения
        self.root.destroy()
