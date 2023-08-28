PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS item (
	itemID integer PRIMARY KEY AUTOINCREMENT,
	itemName varchar,
    itemDescription varchar,
    itemStatus integer
);

CREATE TABLE IF NOT EXISTS list (
	listID integer PRIMARY KEY AUTOINCREMENT,
	listName varchar,
    listDescription varchar,
	listStatus integer,
	createDate date,
	completeDate date
);

CREATE TABLE IF NOT EXISTS listItem (
	listID integer,
	itemID integer,
    CONSTRAINT fk_listid
    FOREIGN KEY(listID) REFERENCES list(listID) ON DELETE CASCADE,
    CONSTRAINT fk_itemid
    FOREIGN KEY(itemID) REFERENCES item(itemID) ON DELETE CASCADE
);