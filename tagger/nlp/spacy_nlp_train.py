import os
import spacy
from spacy.tokens import DocBin


def convert_multi_label(filename):
    db = DocBin()
    nlp = spacy.load('ru_core_news_lg')

    for file in os.listdir('../tags_data'):
        cat_dict = {cat[:-4]: 0 for cat in os.listdir('../tags_data')}
        doc = nlp('\n'.join((open(f'../tags_data/{file}', 'r', encoding='UTF-8').read().split('\n')[1:])))
        doc.cats = cat_dict
        db.add(doc)

    db.to_disk(filename)

convert_multi_label('training_multi_label.spacy')