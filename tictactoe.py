import sys

class Player() :
	def __init__(self,sign,name) :
		self.sign = sign
		self.name = name
	def play(self,case) :
		if case.occ != "#" :
			print("Case is already taken, consider playing something else")
			return(False)
		else :
			case.fill(self.sign)
			return(True)
class Case() :
	def __init__(self):
	 	self.occ = "#"
	def fill(self,sign) :
	 	self.occ = sign

def initcases() :
	caselist = []
	for i in range(9) :
		caselist.append(Case())
	return(caselist)

def iswon(p,cl) :
	for i in range(0,7,3) :
		if cl[0+i].occ == cl[1+i].occ ==cl[2+i].occ == p.sign :
			return(True)
	for i in range(0,3,1) :
		if cl[0+i].occ == cl[3+i].occ == cl[6+i].occ == p.sign :
			return(True)
	if cl[0].occ == cl[4].occ == cl[8].occ == p.sign :
		return(True)
	if cl[2].occ == cl[4].occ == cl[6].occ == p.sign  :
		return(True)
	return(False)

def isdraw(cl) :
	for case in cl :
		if case.occ == "#" :
			return(False)
	return(True)


def display(cl) :
	print(cl[0].occ+"|"+cl[1].occ+"|"+cl[2].occ+"\n-----\n"+cl[3].occ+"|"+cl[4].occ+"|"+cl[5].occ+"\n-----\n"+cl[6].occ+"|"+cl[7].occ+"|"+cl[8].occ+"\n")

def game() :
	Jhon = Player("X","Jhon")
	Jack = Player("O","Jack")
	cl = initcases()
	while True :
		if iswon(Jhon,cl) :
			print("Jhon won the game !")
			sys.exit()
		if iswon(Jack,cl) :
			print("Jack won the game !")
			sys.exit()
		if isdraw(cl) :
			print("This is a draw :(")
			sys.exit()
		display(cl)
		while not Jhon.play(cl[int(input("Your turn, Jhon ! play 123456789\n"))-1]) :
			pass
		print("\n\n\n\n")
		display(cl)
		if iswon(Jhon,cl) :
			print("Jhon won the game !")
			sys.exit()
		if iswon(Jack,cl) :
			print("Jack won the game !")
			sys.exit()
		if isdraw(cl) :
			print("This is a draw :(")
			sys.exit()
		while not Jack.play(cl[int(input("Your turn, Jack ! play 123456789\n"))-1]) :
			pass
		print("\n\n\n\n")

if __name__ == '__main__' :
	game()
