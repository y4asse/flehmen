import functions_framework
import tweepy
import os
from dotenv import load_dotenv

load_dotenv()

dummy_data = {
    "data": [
        {
            "id": 1862851782130692427,
            "text": "https://t.co/zepw65R78X"
        },
        {
            "id": 1862488775877959728,
            "text": "@hizikizikia あーしのことおばちゃんだと思ってんの？ころすよ？"
        },
        {
            "id": 1862433562118824407,
            "text": "春から一人暮らし物件探しに悩んだら、コンビニと薬局！\nどうせそんな自炊しない！\n消耗品はAmazonで買うと普通に割高！送料かかるときもある！"
        },
        {
            "id": 1862420442633678901,
            "text": "お腹痛すぎるしぬど"
        },
        {
            "id": 1862410574539497844,
            "text": "お腹いたいのもう治らない😡‼️\nnこれから絶対下半身不全になって寝たきり😡‼️痛すぎ‼️無理‼️\nってなるのに3日とかでちゃんと治るのオモロ人体"
        },
        {
            "id": 1862406362506174899,
            "text": "お腹痛すぎて死ぬこれ腰砕けるんとちゃうか"
        },
        {
            "id": 1862400573729096129,
            "text": "お腹痛いよ🥲"
        },
        {
            "id": 1862393746333344046,
            "text": "メンタルブレイクにプリンセスウイーク重なって落ち込みプリンセスになってる"
        },
        {
            "id": 1862355824775139707,
            "text": "めちゃくちゃ体重減ったけどこれ絶対筋肉が無くなってるだけやな"
        },
        {
            "id": 1862121088924475544,
            "text": "本気の自炊 https://t.co/l7x7cklUS3"
        }
    ],
    "includes": {},
    "errors": [],
    "meta": {
        "result_count": 10,
        "newest_id": "1862851782130692427",
        "oldest_id": "1862121088924475544",
        "next_token": "7140dibdnow9c7btw4b2e9h1d7yia5b1p89rbrfcvochg"
    }
}

@functions_framework.http
def flehmen_batch(request):
    CONSUMER_KEY = os.getenv("X_API_KEY")
    ACCESS_TOKEN = os.getenv("X_ACCESS_TOKEN")
    BEARER_TOKEN = os.getenv("X_BEARER_TOKEN")
    CONSUMER_SECRET = os.getenv("X_API_KEY_SECRET")
    ACCESS_SECRET = os.getenv("X_ACCESS_TOKEN_SECRET")
    
    client = tweepy.Client(
        consumer_key = CONSUMER_KEY,
        consumer_secret = CONSUMER_SECRET,
        access_token = ACCESS_TOKEN,
        access_token_secret = ACCESS_SECRET,
        bearer_token=BEARER_TOKEN,
    )
    client
    # user_name = "dinsei_"
    
    # user = client.get_user(
    #     username=user_name,
    # )
    
    # data = client.get_users_tweets(
    #     id=user.data.id,
    #     user_auth=True,
    #     max_results=100,
    #     tweet_fields=["created_at"]
    # )
    # print(data)
    
    dummy_data
    
    return "success!"