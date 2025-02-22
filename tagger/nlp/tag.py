from spacy.compat import pickle


def tagger(description):
    file = open(f"./models/tagger.pkl", "rb")
    nlp = pickle.load(file)
    tags_confs = dict(
        sorted(nlp(description).items(), key=lambda item: item[1], reverse=True)
    )

    output = []

    for tag, conf in tags_confs.items():
        if conf >= 0.3:
            output.append(tag)

    return output

    # return str([i for i in tags_confs][:3])
