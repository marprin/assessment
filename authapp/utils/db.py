from pony.orm import Database
from settings import DB_PATH

db = Database()
db.bind(provider="sqlite", filename=DB_PATH, create_db=True)
