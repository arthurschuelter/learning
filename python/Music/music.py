import numpy as np
import math

A = 440.0
notes = ["A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"]

def frequency(semitone, base = A):
    f = math.pow(2, semitone / 12.0) * base
    return f

def main():
    print("Music.py!")
    for i in range(-12, 13):
        print("\t" + notes[i%12], end="")
        f = frequency(i)
        print("%d:\t%.1f Hz" % (4 - math.log(f/A, 0.5), f))

if __name__ == "__main__": main()