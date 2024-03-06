import numpy as np

def plot_move(f, dx, dy):
    return [f[0] + dx, f[1] + dy]

def plot_scale(f, M, dx, dy):
    return [(f[0] - M[0]) * dx + M[0], (f[1] - M[1]) * dy + M[1]]

def plot_rotate(f, M, angle):
    angle_rad = np.deg2rad(angle)
    return [(f[0] - M[0]) * np.cos(angle_rad) - (f[1] - M[1]) * np.sin(angle_rad) + M[0], (f[0] - M[0]) * np.sin(angle_rad) + (f[1] - M[1]) * np.cos(angle_rad) + M[1]]
    
def alt_plot_move(f, dx, dy):
    new_f = []
    for point in f:
        new_f.append((point[0] + dx, point[1] + dy))
    return new_f

def alt_plot_scale(f, M, dx, dy):
    new_f = []
    for point in f:
        new_f.append(((point[0] - M[0]) * dx + M[0], (point[1] - M[1]) * dy + M[1]))
    return new_f

def alt_plot_rotate(f, M, angle):
    angle_rad = np.deg2rad(angle)
    new_f = []
    for point in f:
        new_f.append(((point[0] - M[0]) * np.cos(angle_rad) - (point[1] - M[1]) * np.sin(angle_rad) + M[0],
                      (point[0] - M[0]) * np.sin(angle_rad) + (point[1] - M[1]) * np.cos(angle_rad) + M[1]))
    return new_f
