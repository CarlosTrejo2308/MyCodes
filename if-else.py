#Getting Integer
n = int(input())

#Is odd?
if ( (n % 2) != 0 ):
    print("Weird")
    
#Even
elif ( n <= 5):
    print("Not Weird")
elif ( n <= 20 ):
    print("Weird")
else:
    print("Not Weird")

