import spacy


def tag(description):
    nlp = spacy.load('../models/textcat_multilabel_model/model-last')
    tags_confs = dict(
        sorted(nlp(description).cats.items(), key=lambda item: item[1], reverse=True)
    )

    output = []
    mid = 0

    for conf in tags_confs.items():
        mid += conf[1]

    mid /= len(tags_confs)

    for tag, conf in tags_confs.items():
        if conf >= mid:
            output.append(tag)

    return output