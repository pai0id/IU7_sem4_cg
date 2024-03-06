import tkinter as tk  # Импорт модуля для создания графического интерфейса
from geometry_app import GeometryApp  # Импортируем класс GeometryApp из модуля geometry_app

if __name__ == "__main__":
    root = tk.Tk()  # Создаем основное окно tkinter
    app = GeometryApp(root)  # Создаем экземпляр класса GeometryApp, передавая основное окно как родительский элемент
    root.mainloop()  # Запускаем главный цикл обработки событий tkinter, чтобы приложение оставалось активным
