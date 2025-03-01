# Rural Potatoes

Rural Potatoes - это ваш главный помощник в поиске фильмов на вечер. Он позволит вам не только найти новые впечатления, но и поделиться ими с другими людьми!
В решении реализованы:


- **Комментарии**.
- **Лайки** и **дизлайки** к фильмам.
- Поиск **схожих** по вашему вводу тегов.
- Система **рекомендаций**.
- **Готовая** база данных, содержащая в себе **около 700** фильмов.
- **Аунтификация** пользователя.

<p align="center">
  <img src="assets/video_record_1.gif" width="738">
</p>

## Установка 
```shell
git clone https://github.com/bebrochkas/rural_potatoes.git
cd rural_potatoes
```

## Использование 
Скачать модель [отсюда](/releases) в `/tagger/models/` для функционала умного поиска в реальном времени

### Запуск
Убедитесь что Docker уставновлен на вашей системе

В корневой папке репозитория запустите контейнер:
```shell
docker compose up
```

## Дополниетльно
### При желеании можете обучть модель самостоятельно
```shell
cd tagger/nlp
python -m spacy train spacy_categorization_pipeline.cfg
--paths.train training_multi_label.spacy 
--paths.dev test_multi_label.spacy 
--output ../models/textcat_multilabel_model
```
и дождаться окончания процесса обучения обучения модели.

### Интерфейс будет доступен на [localhost:3000](http://loclhost:3000)🎉
