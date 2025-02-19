import os
from classy_classification import ClassyClassifier
from spacy.compat import pickle


def learn_model():
    data = {''.join(list(i)[:-4]):open(f'{os.getcwd()}\\tags_data\\{i}', 'r', encoding='UTF-8').read().split('\n')[1:] for i in os.listdir(f'{os.getcwd()}\\tags_data')}

    nlp = ClassyClassifier(
        data= data,
        model= "sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2",
        multi_label= True,
    )

    with open(f'{os.getcwd()}\\models\\tagger.pkl', 'wb') as file:
       pickle.dump(nlp, file)