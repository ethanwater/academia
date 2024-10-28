import requests
from bs4 import BeautifulSoup


def fetch_googledoc_input_data(googledoc_url):
    export_url = googledoc_url.replace('/edit', '/export?format=html')
    response = requests.get(export_url)

    if response.status_code != 200:
        print("[ERROR] Unable to retrieve data from provided URL")
        return
   
    soup = BeautifulSoup(response.text, 'html.parser')

    grid = soup.find('table')
    if not grid:
        print("[ERROR] Grid unavailable")
        return

    grid_data = []

    for row in grid.find_all('tr')[1:]: 
        row_data = [cell.get_text(strip=True) for cell in row.find_all('td')]
        grid_data.append(row_data)
    return grid_data

def convert_to_int_helper(grid):
    for row in grid:
        for i, cell in enumerate(row):
            if isinstance(cell, str) and cell.isdigit():
                row[i] = int(cell)

def decode_secret_message(grid):
    convert_to_int_helper(grid)
    max_x = max(row[0] for row in grid)  
    max_y = max(row[2] for row in grid)  

    matrix = [[' ' for _ in range(max_x + 1)] for _ in range(max_y + 1)]
    
    for row in grid:
        column_idx = row[0]
        row_idx = row[2]
        matrix[row_idx][column_idx] = row[1]

    matrix.reverse()
    return matrix

def reveal_secret_message(googledoc_url):
    googledoc_grid_data = fetch_googledoc_input_data(googledoc_url)
    secret_message = decode_secret_message(googledoc_grid_data)
    for row in secret_message:
        print("".join(row).rstrip())

if __name__ == '__main__':
    test_input_url = "https://docs.google.com/document/d/e/2PACX-1vQGUck9HIFCyezsrBSnmENk5ieJuYwpt7YHYEzeNJkIb9OSDdx-ov2nRNReKQyey-cwJOoEKUhLmN9z/pub"
    reveal_secret_message(test_input_url)

