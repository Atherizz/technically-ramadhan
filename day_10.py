def make_readable(seconds: int) -> str:
    hours = seconds // 3600
    minute = seconds % 3600 // 60
    second = seconds % 60
    
    return f"{hours:02}:{minute:02}:{second:02}"
    
print(make_readable(70))
    
    
    
    