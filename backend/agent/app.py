from flask import Flask, jsonify, request
from flask_cors import CORS
from dotenv import load_dotenv
from browser_use import Agent, Browser, BrowserConfig, SystemPrompt
from langchain_openai import ChatOpenAI
import json
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
# planner_llm = ChatOpenAI(model='gpt-4o')

@app.route('/', methods=['POST'])
async def flask_app():
    body = request.get_data()
    body = json.loads(body)
    x_username = body['x_username']
    print(x_username)
    
    print(body)
    initial_actions = [
        {'open_tab': {'url': 'https://twitter.com'}},
        {'scroll_down': {'amount': 0}},
    ]

    agent = Agent(
		task=f"""{x_username} のidのアカウントページにアクセスして，そのアカウントの最新のツイートを取得してください．""",
		llm=ChatOpenAI(model='gpt-4o'),
        browser=Browser(config=BrowserConfig(
            headless=True
        )),
        sensitive_data=sensitive_data,
        initial_actions=initial_actions,
        system_prompt_class=MySystemPrompt,
        max_failures=1,
    )
    result = await agent.run()
    final_result = result.final_result()
    print(final_result)
    return jsonify({"result": final_result})

@app.route('/ok')
async def ok():
    return jsonify({'message': 'ok'})


if __name__ == '__main__':
    app.run(host='0.0.0.0')