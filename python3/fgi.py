stuff = {'rope': 1, 'torch': 6, 'gold coin': 42, 'dagger': 1, 'arrow': 12}

def displayInventory(inventory):
    print("Inventory:")
    item_total = 0
    for k, v in inventory.items():
        print(str(v) + ' ' + k)
        item_total += v
    print("Total number of items: " + str(item_total))


def addToInventory(inventory, addedItems):
    for i in addedItems:
        inventory[i] = inventory.setdefault(i, 0) + 1
    return inventory
    # your code goes here

inv = {'gold coin': 42, 'rope': 1}
dragonLoot = ['gold coin', 'dagger', 'gold coin', 'gold coin', 'ruby']
print("-------Before loot-------")
displayInventory(inv)
inv = addToInventory(inv, dragonLoot)
print("-------After loot-------")
displayInventory(inv)

# displayInventory(stuff)
