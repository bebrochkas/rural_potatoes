import os
from classy_classification import ClassyClassifier
from spacy.compat import pickle


def train_model():
    data = {
        "".join(list(i)[:-4]): open(f"../tags_data/{i}", "r", encoding="UTF-8")
        .read()
        .split("\n")[1:]
        for i in os.listdir("../tags_data")
    }

    nlp = ClassyClassifier(
        data=data,
        multi_label=True,
        model="cointegrated/rubert-tiny2",
    )

    os.makedirs("../models", exist_ok=True)

    with open("../models/tagger.pkl", "wb") as file:
        pickle.dump(nlp, file)