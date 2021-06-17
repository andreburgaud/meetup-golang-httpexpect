from typing import Optional
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import uvicorn

app = FastAPI()


db = {}


class Item(BaseModel):
    name: str
    author: str
    price: float


@app.get("/hi")
async def root():
    return {"message": "Hello Gophers!"}


@app.get("/ids/{item_id}")
async def read_item(item_id):
    return {"item_id": item_id}


@app.post("/items", response_model=Item)
async def create_item(item: Item):
    db[item.name] = item
    return item


@app.get("/items/{name}")
async def get_item(name):
    if name in db:
        return db[name]
    else:
        raise HTTPException(status_code=404, detail="Item not found")


@app.get("/items")
async def get_all_items():
    if db:
        return list(db.values())
    else:
        raise HTTPException(status_code=404, detail="Item not found")

@app.get("/clear")
async def delete_items():  # For testing
    db.clear()
    return {"action": "db cleared"}


if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=9999)