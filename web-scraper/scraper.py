import pandas as pd
import requests
from bs4 import BeautifulSoup
from bs4.element import Tag
from sqlalchemy import create_engine


def main():
    categories = ['batting', 'pitching', 'fielding']
    for category in categories:
        url = f'https://finland.wbsc.org/en/events/2024-superbaseball/stats/general/team/31569/all/{category}'
        df = scrape(url)
        write_to_database(category, df, 'sqlite:///superbaseball_stats.db')


def scrape(url: str) -> pd.DataFrame:
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    table = soup.find('table')

    return extract_from_table(table)


def extract_from_table(table: Tag) -> pd.DataFrame:
    headers = []
    thead = table.find('thead')
    for th in thead.find_all('th'):
        headers.append(th.text.strip())

    rows = []
    tbody = table.find('tbody')
    for tr in tbody.find_all('tr'):
        row = []
        for td in tr.find_all('td'):
            row.append(td.text.strip())
        rows.append(row)

    return pd.DataFrame(rows, columns=headers)


def write_to_database(table_name: str, df: pd.DataFrame, database_url: str) -> None:
    engine = create_engine(database_url)
    df.to_sql(table_name, con=engine, if_exists='replace', index=False)


if __name__ == '__main__':
    main()
