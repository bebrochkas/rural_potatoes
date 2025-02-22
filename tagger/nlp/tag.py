from spacy.compat import pickle


def tagger(description):
    file = open(f"./models/tagger.pkl", "rb")
    nlp = pickle.load(file)
    tags_chance = dict(sorted(nlp(description).items(), key=lambda item: item[1], reverse=True))

    return [i for i in tags_chance][:3]