import os
from spacy.compat import pickle


def tagger(description):
    file = open(f"{os.getcwd()}\\models\\tagger.pkl", "rb")
    nlp = pickle.load(file)

    return nlp(description)