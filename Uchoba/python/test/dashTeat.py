import dash
from dash import html, dcc

# Создаем приложение
app = dash.Dash(__name__)

# Определяем макет (layout) - что будет отображаться
app.layout = html.Div([
    html.H1("Basic Dash Dashboard"),

    dcc.Graph(
        id="sales-chart",
        figure={
            "data": [
                {
                    "x": [1, 2, 3, 4, 5],
                    "y": [5, 4, 7, 4, 8],
                    "type": "line",
                    "name": "Trucks"
                },
                {
                    "x": [1, 2, 3, 4, 5],
                    "y": [6, 3, 5, 3, 7],
                    "type": "bar",
                    "name": "Ships"
                }
            ],
            "layout": {
                "title": "Basic Dashboard"
            }
        }
    )
])

# Запускаем сервер
if __name__ == "__main__":
    app.run(debug=True)