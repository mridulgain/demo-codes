#! python3
   # backupToZip.py - Copies an entire folder and its contents into
   # a ZIP file whose filename increments.

import zipfile, os
import pyinputplus as pyip

def backupToZip(target):
	# Back up the entire contents of "folder" into a ZIP file.
	zipFilename =  target + '.zip'
	if os.path.exists(zipFilename):
		response = pyip.inputYesNo("Zip file already exists, Would you like to override? ")
		if response == 'no':
			return
		else:
			print(f'Creating {zipFilename}...')
			backupZip = zipfile.ZipFile(zipFilename, 'w')
	print(f'Creating {zipFilename}...')
	backupZip = zipfile.ZipFile(zipFilename, 'w')
	if os.path.isdir(target):
		for foldername, subfolders, filenames in os.walk(target):
			# print(f'Adding files in {foldername}...')
			backupZip.write(foldername)
			for filename in filenames:
				print(f'Adding {os.path.join(foldername,filename)}...')
				backupZip.write(os.path.join(foldername,filename))
	else:
		backupZip.write(target)
	backupZip.close()
	print('Done.')

backupToZip('test')