def alt_or(lst):
    if not lst:
        return None
    else:
        if True in lst:
            return True
        else:
            return False
