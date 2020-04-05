import datetime
newyear = False
secret_message = "¡Feliz Año Nuevo!"

while (not newyear):
	year = datetime.date.today().year

	if (year == 2020):
		newyear = True

print(secret_message)