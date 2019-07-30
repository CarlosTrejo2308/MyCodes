#helps the determine the number of a fib pos
#@param l2: a list containing 3 elements of the fib seq [cur - 2, cur - 1, cur]
#@param tar: the desired position to get
#@param cur: the current position calculated
#@return the number of that given position in the fib seq
def fibhelp(l2, tar, cur):
        #some error occured, current exceeded target
        if (tar < cur):
            return -1
        #target found, return it
	if(cur == tar):
		return l2[2]
	    
	#calculate next number:
	#shift left 1 position and calculate the most right number (cur + 1) by adding the 2 previous numbers
	# [cur - 2, cur - 1, cur] --> [cur - 1, cur - 2, cur + 1]
	else:
		#shift
		l2[0] = l2[1]
		l2[1] = l2[2]
		#calculate cur + 1
		l2[2] = l2[0] + l2[1]
		#recursevly use fibhelp but with the next current 
		return fibhelp(l2, tar, cur + 1)

#main fibonacci function
#returns the resulting number in acordance of the position given on a fibonacci sequence
# @param num: the position of the fib (num > 0)
# @return the number of that given position in the fib seq
def fib(num):
        #error, there´s no non-positive position
        if (num < 0):
            return -1
        #first pos
	elif (num == 1):
		return 0
	#secnond pos
	elif (num == 2):
		return 1
	#rest of the positions, use fib helper to calculate the number
	else:
		return fibhelp([0,1,1], num, 3)
		

#main function
def main():

    #asks for input and prints rhe result pos in the fib sequence
    x = int(input())
    print(fib(x))

	
main()
