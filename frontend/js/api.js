async function getAllLists() {
    let response = await fetch("http://localhost:8080/lists");
    let shopLists = await response.json();
    return shopLists
}

async function getListItems(listID) {
    let response = await fetch(`http://localhost:8080/listitems/${listID}`)
    let listItems = await response.json()
    return listItems
}

async function saveNewList(list){
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(list)
    }
    let response = await fetch("http://localhost:8080/newlist", options);
    let rlist = await response.json();
    return rlist;
}

async function saveNewItem(item){
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(item)
    }
    let response = await fetch("http://localhost:8080/newitem", options);
    let ritem = await response.json();
    return ritem;
}

async function updateItem(item) {
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(item)
    }
    let response = await fetch("http://localhost:8080/updateitem", options);
    let ritem = await response.json();
    return ritem;
}

async function deleteItem(item) {
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(item)
    }
    let response = await fetch("http://localhost:8080/deleteitem", options);
    let resp = await response.text();
    return resp;
}
async function setitemCompleted(item) {
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(item)
    }
    let response = await fetch("http://localhost:8080/completeitem", options);
    let resp = await response.text();
    return resp;
}


async function updateList(list){
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(list)
    }
    let response = await fetch("http://localhost:8080/updatelist", options);
    let resp = await response.text();
    return resp;
}


async function deleteList(list){
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(list)
    }
    let response = await fetch("http://localhost:8080/deletelist", options);
    let resp = await response.text();
    return resp;
}

async function setListCompleted(list){
    let options = {
        method: 'POST',
        headers: {
            'Content-Type': 
                'application/json;charset=utf-8'
        },
        body: JSON.stringify(list)
    }
    let response = await fetch("http://localhost:8080/completelist", options);
    let resp = await response.text();
    return resp;
}