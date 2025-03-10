
def alphabet_position(text: str) -> str:
    res = ""
    for letter in text.lower():
        if 'a' <= letter <= 'z':
            char = ord(letter) - ord('a') + 1
            res += str(char) + " "

    return res.strip()

print(alphabet_position("The sunset sets at twelve o' clock."))