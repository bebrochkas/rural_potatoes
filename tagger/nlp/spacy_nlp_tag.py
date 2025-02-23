import spacy


def tag(description):
    nlp = spacy.load('../models/textcat_multilabel_model/model-last')
    tags_confs = dict(
        sorted(nlp(description).cats.items(), key=lambda item: item[1], reverse=True)
    )

    output = []

    for tag, conf in tags_confs.items():
        if conf >= 0.3:
            output.append(tag)

    return output