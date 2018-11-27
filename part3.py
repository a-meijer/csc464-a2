'''
NAME: part3.py
AUTHOR: Andrew Meijer V00805554
INFORMATION:
This program is an introduction to my project.
The idea for my final project is to verify the correctness of recorded chess
games by using a concurrent approach. In effect, my final project will replay
any chess game, and some threads will read ahead and verify moves before they
have been replayed.
This is relevant to the CAP Theorem because the solution must be linearizably
consistent.
In this program I will begin to develop my project by verifying only the opening
moves. To test the program, a user must use the command line to enter a chess
move for the white pieces such as "e4" or "Nf3."
Ideally, the program will test if the move is valid according to the game state
and then update the state in a critical section.
Since this sample program just operates on the starting position, and also
because of time constraints, I do not include a game state.
'''

#!/usr/bin/python3
import threading
import time
import re

move = ""
done = 0

class freethread(threading.Thread):
   def run(self):
       print ("Welcome to chess.")
       global move
       global done

       m = input("Using the white pieces, enter a valid starting move: ")

       # Verify the input is proper syntax.
       list = ["a3","a4","b3","b4","c3","c4","d3","d4","e3","e4","f3","f4","g3","g4","h3","h4","Na3","Nc3","Nf3","Nh3"]

       for element in list:
           if(re.match(element, m)):
               print("Success.")
               done = 1

       if(done == 0):
           print("Not a legal move.")
           done = 1



# Create new threads
t1 = freethread()
# Threads can only be started once
t1.start()

# wait
while(done == 0):
    pass
print ("Exiting Main Thread.")
