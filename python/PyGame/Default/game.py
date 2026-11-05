import pygame as pg
import random
##
from Settings import Settings
##

def setup():
	# resolution = (1280, 720)
	resolution = (960, 720)

	pg.init()

	settings = Settings(resolution, 60)
	pg.display.set_caption('Joguito')

	screen  = pg.display.set_mode(resolution)

	settings.font = pg.font.Font(None, 30)
	settings.clock = pg.time.Clock()

	return screen, settings

def chess(screen, settings):
	color = (180, 180, 255)
	black = (30,30,30)
	white = (220,220,220)

	w = settings.resolution[0]
	h = settings.resolution[1]

	num = 8
	thicc = 10
	cont = 0

	for y in range(0, num):
		for x in range(0, num):

			if(cont % 2 == 0):
				color = white
			else:
				color = black

			pg.draw.rect(screen, color,
							( ((w-h)/2) + (h * x) / num, 	# X Starting point
							(h * y) / num,	 				# Y Starting point
							(h / num),		 				# X size
							(h / num)))		 				# Y size
			cont += 1
		cont +=1

def draw(screen, settings):
	background = (0, 0, 0)

	screen.fill(background)
	chess(screen, settings)
	update(screen, settings)


def update(screen, settings):
	pg.draw.rect(screen, pg.Color(0, 0, 0, 50), (50, 50, 75, 20))
	fps = settings.font.render("FPS: " + str(int(settings.clock.get_fps())), True, pg.Color('white'))
	screen.blit(fps, (50, 50))
	pg.display.flip()
	settings.clock.tick(settings.framerate)

def main():
	print("Starting...")

	screen, settings = setup()

	running = True
	while(running):
		draw(screen, settings)
		for event in pg.event.get():
			if event.type == pg.QUIT:
				running = False

if __name__ == "__main__": main()
