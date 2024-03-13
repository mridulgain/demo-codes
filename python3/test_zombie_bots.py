import zombiedice
import random

class MyZombie:
    def __init__(self, name):
        # All zombies must have a name:
        self.name = name

    def turn(self, gameState):
        # gameState is a dict with info about the current state of the game.
        # You can choose to ignore it in your code.

        diceRollResults = zombiedice.roll() # first roll
        # roll() returns a dictionary with keys 'brains', 'shotgun', and
        # 'footsteps' with how many rolls of each type there were.
        # The 'rolls' key is a list of (color, icon) tuples with the
        # exact roll result information.
        # Example of a roll() return value:
        # {'brains': 1, 'footsteps': 1, 'shotgun': 1,
        #  'rolls': [('yellow', 'brains'), ('red', 'footsteps'),
        #            ('green', 'shotgun')]}

        # REPLACE THIS ZOMBIE CODE WITH YOUR OWN:
        brains = 0
        while diceRollResults is not None:
            brains += diceRollResults['brains']

            if brains < 2:
                diceRollResults = zombiedice.roll() # roll again
            else:
                break
class Bot1:
    def __init__(self, name):
        self.name = name

    def turn(self, gameState):
        diceRollResults = zombiedice.roll() # first roll
        options = [True, False]
        if random.choice(options) == True:
            diceRollResults = zombiedice.roll()
        else:
            return

class Bot4:
    def __init__(self, name):
        self.name = name

    def turn(self, gameState):
        sg = 0
        for i in range(4):
            diceRollResults = zombiedice.roll()
            sg += diceRollResults['shotgun']
            if sg >= 2:
                break

class Bot5:
    def __init__(self, name):
        self.name = name

    def turn(self, gameState):
        diceRollResults = zombiedice.roll()
        sg = diceRollResults['shotgun']
        br = diceRollResults['brains']
        while diceRollResults is not None:
            if sg > br:
                break
            diceRollResults = zombiedice.roll()

zombies = (
    zombiedice.examples.RandomCoinFlipZombie(name='Random'),
    zombiedice.examples.RollsUntilInTheLeadZombie(name='Until Leading'),
    zombiedice.examples.MinNumShotgunsThenStopsZombie(name='Stop at 2 Shotguns', minShotguns=2),
    zombiedice.examples.MinNumShotgunsThenStopsZombie(name='Stop at 1 Shotgun', minShotguns=1),
    MyZombie(name="book example"),
    # Add any other zombie players here.
    Bot1(name='bot1'),
    Bot4('Bot4'),
    Bot5('Bot5')
)

# Uncomment one of the following lines to run in CLI or Web GUI mode:
zombiedice.runTournament(zombies=zombies, numGames=100)
#zombiedice.runWebGui(zombies=zombies, numGames=1000)
