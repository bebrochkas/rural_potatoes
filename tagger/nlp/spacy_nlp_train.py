import os
import spacy
from spacy.tokens import DocBin


def convert_multi_label(filename_train, filename_test):
    db_train = DocBin()
    nlp = spacy.load('ru_core_news_lg')

    for file in os.listdir('../tags_data'):
        cat_dict = {cat[:-4]: 0 for cat in os.listdir('../tags_data')}
        doc = nlp('\n'.join((open(f'../tags_data/{file}', 'r', encoding='UTF-8').read().split('\n')[1:])))
        doc.cats = cat_dict
        db_train.add(doc)

    db_train.to_disk(filename_train)
    db_test = DocBin()

    for file in os.listdir('../test_tags_data'):
        cat_dict = {cat[:-4]: 0 for cat in os.listdir('../tags_data')}
        doc = nlp('\n'.join((open(f'../test_tags_data/{file}', 'r', encoding='UTF-8').read().split('\n')[1:])))
        doc.cats = cat_dict
        db_test.add(doc)

    db_test.to_disk(filename_test)

convert_multi_label('training_multi_label.spacy', 'test_multi_label.spacy')