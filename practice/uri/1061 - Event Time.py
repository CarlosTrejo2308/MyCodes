entrada1 = input();
entrada2= input();
entrada3 = input();
entrada4= input();

dia1= int(entrada1[4:])
dia2= int(entrada3[4:])

hora1= int(entrada2[0:2])
min1= int(entrada2[5:7])
seg1= int(entrada2[10:])

hora2= int(entrada4[0:2])
min2= int(entrada4[5:7])
seg2= int(entrada4[10:])

dias = int(dia2-dia1-1)
horas= int(24-hora1+hora2)

if horas>=24:
    dias= int(dias+1)
    horas= int(horas-24)
minut= int((60-min1)-(60-min2))
if minut<0:
    horas=horas-1
    minut=60+minut
segu= int((60-seg1)-(60-seg2))
if segu<0:
    minut=minut-1
    segu=60+segu

dias = str(dias)
horas= str(horas)
minut= str(minut)
segu= str(segu)

print(dias + " dia(s)")
print(horas + " hora(s)")
print(minut + " minuto(s)")
print(segu + " segundo(s)")
