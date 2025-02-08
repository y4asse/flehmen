from flask import Flask, jsonify
from flask_cors import CORS
from dotenv import load_dotenv
from browser_use import Agent, Browser, BrowserConfig, SystemPrompt
from langchain_openai import ChatOpenAI

class MySystemPrompt(SystemPrompt):
    def important_rules(self) -> str:
        # Get existing rules from parent class
        existing_rules = super().important_rules()

        # Add your custom rules
        new_rules = """
        答えは日本語で出力してください．
"""

        # Make sure to use this pattern otherwise the exiting rules will be lost
        return f'{existing_rules}\n{new_rules}'


load_dotenv()

app = Flask(__name__)
CORS(app)

sensitive_data = {'x_user_name': 'test_ai_aya', 'x_password': 'Yama0207'}
planner_llm = ChatOpenAI(model='gpt-4o')

@app.route('/')
async def flask_app():
    initial_actions = [
        {'open_tab': {'url': 'https://twitter.com/nyantarou_030'}},
        {'scroll_down': {'amount': 0}},
    ]

    agent = Agent(
		task="""https://twitter.com/nyantarou_030 にアクセスして，そのアカウントの新年最初のツイートを取得してください．""",
		llm=ChatOpenAI(model='gpt-4o'),
        browser=Browser(config=BrowserConfig(
            headless=True
        )),
        sensitive_data=sensitive_data,
        planner_llm=planner_llm,
        initial_actions=initial_actions,
        system_prompt_class=MySystemPrompt
    )
    result = await agent.run()
    final_result = result.final_result()
    print(final_result)
    # resultをjson形式で返す
    return jsonify({"result": final_result})

    return final_result

if __name__ == '__main__':
    app.run(host='0.0.0.0')