from flask import Flask
from flask_cors import CORS
from dotenv import load_dotenv
from browser_use import Agent
from langchain_openai import ChatOpenAI

load_dotenv()

app = Flask(__name__)
CORS(app)

@app.route('/')
def flask_app():
    agent = Agent(
		task="""@@nyantarou_030 のTwitter idのアカウントの直近の10ツイートを取得してください．
                ログイン情報はuser_id=test_ai_aya, password=Yama0207です．""",
		llm=ChatOpenAI(model='gpt-4o'),
	)
    result = agent.run()
    print(result)
    return "Hello, World!"

if __name__ == '__main__':
    app.run(host='0.0.0.0')