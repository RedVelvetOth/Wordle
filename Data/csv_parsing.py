# Cleans the data and generates a txt file of the 5 letters words and their defintion 

import pandas as pd


data = pd.read_csv("Data/OPTED-Dictionary.csv").convert_dtypes()

data.drop_duplicates("Word",inplace=True)
data.drop("POS",inplace=True,axis=1)
data = data[data["Count"] == "5"]
data = data[data["Word"].str.contains(' ') == False]
data.reset_index(inplace=True)
data.drop({"Count","index"},inplace=True,axis=1)
data["Word"] = data["Word"].str.lower()

data.to_csv("Data/dictionnary.txt",index=False,header=False)