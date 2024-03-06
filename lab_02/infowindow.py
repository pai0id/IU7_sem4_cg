import tkinter as tk  # Импорт модуля для создания графического интерфейса

# Вывод окна с информацией
class InfoWindow:
    def __init__(self, root, string):
        self.root = tk.Toplevel(root)
        self.root.title("Info")
        self.root.geometry(f"{int(root.winfo_screenwidth() * 0.3)}x{int(root.winfo_screenheight() * 0.2)}+1100+200")
        
        self.text = tk.Text(self.root, font="system", wrap=tk.WORD, height=7)
        self.text.pack(padx=10, pady=10, fill=tk.X)
        self.text.insert(tk.END, string)
        self.text.configure(state=tk.DISABLED)
        
        tk.Button(self.root, text="Выход", command=lambda:self.root.destroy()).pack(padx=10, pady=10)