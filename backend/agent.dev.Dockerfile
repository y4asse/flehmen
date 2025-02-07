# Pythonランタイムを親イメージとして使用
FROM python:3.13.0

# 作業ディレクトリを/appに設定
WORKDIR /app

# 現在のディレクトリの内容をコンテナ内の/appにコピー
COPY ./agent .

# requirements.txtで指定された必要なパッケージをインストール
RUN pip install -r requirements.txt
RUN pip install pipreqs

# ポートの公開
EXPOSE 5000
CMD ["flask","--app", "app.py", "run", "--host=0.0.0.0", "--debug"]