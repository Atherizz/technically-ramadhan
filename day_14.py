from typing import List

def luck_check(str_num: str) -> bool:
    if not str_num:
        raise ValueError("Empty string is not valid")
    
    for char in str_num:
        if not char.isdigit():
            raise ValueError("Input should contain digits only")
    
    digits = []
    for digit in str_num:
        digits.append(int(digit))
    
    left, right = [], []
    mid = len(digits) // 2
    
    if len(digits) % 2 == 0:
        left = digits[:mid]
        right = digits[mid:]
    else:
        left = digits[:mid]
        right = digits[mid+1:]
    
    return sum_arr(left, 0, len(left) - 1) == sum_arr(right, 0, len(right) - 1)


def sum_arr(arr: List[int], l: int, r: int) -> int:
    if l == r:
        return arr[l]
    else:
        mid = (l + r) // 2
        left_sum = sum_arr(arr, l, mid)
        right_sum = sum_arr(arr, mid + 1, r)
        return left_sum + right_sum