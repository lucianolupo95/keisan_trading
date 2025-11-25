#!/usr/bin/env python3
"""
Katon (Fuego) - módulo de features y transformaciones
"""

def ping():
    """Retorna un mensaje de validación"""
    return "Katon OK"

def simple_feature(x):
    """Función simple: multiplica por 2"""
    return x * 2

if __name__ == "__main__":
    print("Katon activo")
    response = ping()
    print(f"Ping response: {response}")

    # Test simple_feature
    test_value = 5
    result = simple_feature(test_value)
    print(f"simple_feature({test_value}) = {result}")
