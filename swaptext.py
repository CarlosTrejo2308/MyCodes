def swap_text(s):
    result = ""

    for letter in s:
        if letter.isupper():
            result += chr( ord(letter) + 32 )
        elif letter.islower():
            result += chr( ord(letter) - 32 )
        else:
            result += letter

    return result

print( swap_text("hola TODOS") )

print( " ".isupper() ) 