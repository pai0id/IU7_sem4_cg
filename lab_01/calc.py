import itertools  # Импорт модуля для работы с итераторами
import math       # Импорт модуля для математических вычислений
from math import fabs  # Импорт функции fabs из модуля math

def float_eq(a, b):
    # Функция для сравнения двух чисел с плавающей точкой
    return True if fabs(a - b) < 0.00001 else False

def distance(p1, p2):
    # Функция для вычисления расстояния между двумя точками в пространстве
    return math.sqrt((p1[0] - p2[0])**2 + (p1[1] - p2[1])**2)

def triangle_area(p1, p2, p3):
    # Функция для вычисления площади треугольника по его вершинам
    a = distance(p1, p2)
    b = distance(p2, p3)
    c = distance(p3, p1)
    s = (a + b + c) / 2
    return math.sqrt(s * (s - a) * (s - b) * (s - c))

def circumcircle_radius(p1, p2, p3):
    # Функция для вычисления радиуса описанной окружности треугольника
    a = distance(p1, p2)
    b = distance(p2, p3)
    c = distance(p3, p1)
    return (a * b * c) / (4 * triangle_area(p1, p2, p3))

def incircle_radius(p1, p2, p3):
    # Функция для вычисления радиуса вписанной окружности треугольника
    a = distance(p1, p2)
    b = distance(p2, p3)
    c = distance(p3, p1)
    return triangle_area(p1, p2, p3) / (0.5 * (a + b + c))

def circumcircle_center(p1, p2, p3):
    # Функция для вычисления центра описанной окружности треугольника
    x1, y1 = p1
    x2, y2 = p2
    x3, y3 = p3
    D = 2 * (x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2))
    Ux = ((x1**2 + y1**2) * (y2 - y3) + (x2**2 + y2**2) * (y3 - y1) + (x3**2 + y3**2) * (y1 - y2))
    Uy = ((x1**2 + y1**2) * (x3 - x2) + (x2**2 + y2**2) * (x1 - x3) + (x3**2 + y3**2) * (x2 - x1))
    return (Ux / D, Uy / D)

def incircle_center(p1, p2, p3):
    # Функция для вычисления центра вписанной окружности треугольника
    a = distance(p2, p3)
    b = distance(p3, p1)
    c = distance(p1, p2)
    x = (a * p1[0] + b * p2[0] + c * p3[0]) / (a + b + c)
    y = (a * p1[1] + b * p2[1] + c * p3[1]) / (a + b + c)
    return (x, y)

def find_circumcircle(points):
    # Функция для нахождения описанной окружности вокруг треугольника
    p1, p2, p3 = points
    return (circumcircle_center(*points), circumcircle_radius(*points))

def find_incircle(points):
    # Функция для нахождения вписанной окружности в треугольник
    return (incircle_center(*points), incircle_radius(*points))

def find_triangle(points):
    # Функция для нахождения треугольника с наиб
    max_diff = float('-inf')
    max_triangle = None
    max_c_radius = 0
    min_i_radius = 0

    for triangle in itertools.combinations(points, 3):
        if (triangle_area(*triangle) == 0):
            continue
        c_radius = circumcircle_radius(*triangle)
        i_radius = incircle_radius(*triangle)
        diff = c_radius - i_radius
        if diff > max_diff:
            max_diff = diff
            max_triangle = triangle
            max_c_radius = c_radius
            min_i_radius = i_radius

    return max_triangle, math.pi * max_c_radius ** 2, math.pi * min_i_radius ** 2