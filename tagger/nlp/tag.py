from spacy.compat import pickle


def tagger(description):
    file = open(f"./models/tagger.pkl", "rb")
    nlp = pickle.load(file)
    tags_confs = dict(
        sorted(nlp(description).items(), key=lambda item: item[1], reverse=True)
    )

    output = []
    mid = 0

    for conf in tags_confs.items():
        mid += conf[1]

    mid /= len(tags_confs)
    mid += 0.1

    for tag, conf in tags_confs.items():
        if conf >= mid:
            output.append(tag)

    return output