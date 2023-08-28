window.addEventListener("load", function() {
    initialLoad()
});

// GLOBAL LIST and ITEMS ARRAY
let LISTS = []
let ITEMS = []

// fetches the shopping lists and prints them to page.
async function initialLoad(){
    LISTS = await getAllLists()
    if (LISTS == null) {return appendListRow("<span class='panel'>No lists found, consider adding one</span>")}
    for (let i = 0; i < LISTS.length; i++) {
        const row = LISTS[i];
        appendListRow(newListRow(row, i+1))
    }
}



function handleListClick(listID, listnum) {
    //highligth
    //TODO

    // place in config panel.
    fillListForm(listnum-1)

    // Remove old
    clearItems()
    // load items
    loadItems(listID)
}

function handleItemClick(itemIndx) {
    //highligth
    //TODO

    // place in config panel.
    fillItemForm(itemIndx)
}


async function loadItems(listID) {
    ITEMS = await getListItems(listID);
    if (ITEMS == null) {return appendItemRow("<span class='panel'>No items in the selected list</span>")}
    
    for (let i = 0; i < ITEMS.length; i++) {
        const ele = ITEMS[i];
        appendItemRow(newItemRow(ele, i))
    }
}


async function createAndSaveList() {
    let newl = {};
    newl["name"] = document.getElementById("formListName").value;
    newl["description"] = document.getElementById("formListDesc").value;
    newl["status"] = 0;
    let retList = await saveNewList(newl);
    if (LISTS == null) {
        LISTS = []
        clearLists()
    }
    LISTS.push(retList)
    appendListRow(newListRow(retList, LISTS.length))
}

async function createAndSaveItem(){
    let newi = {}
    newi["listID"] = document.getElementById("formListID").valueAsNumber
    newi["name"] = document.getElementById("formItemName").value
    newi["description"] = document.getElementById("formItemDesc").value
    newi["status"] = 0

    let retIt = await saveNewItem(newi)
    if (ITEMS == null) {
        ITEMS = []
        clearItems()
    }
    ITEMS.push(retIt)
    appendItemRow(newItemRow(retIt, ITEMS.length-1))
}

async function handleItemUpdate() {
    let it = ITEMS[document.getElementById("formItemIndx").valueAsNumber]
    it["name"] = document.getElementById("formItemName").value
    it["description"] = document.getElementById("formItemDesc").value
    let r = await updateItem(it)
    console.log(r)
    clearItems()
    loadItems(it["listID"])
}

async function handleDeleteItem(itemIndx) {
    let it = ITEMS[itemIndx];
    let lid = it["listID"];
    let r = await deleteItem(it);
    console.log(r);
    clearItems();
    loadItems(lid);
}

async function handleUpdateList() {
    let li = LISTS[document.getElementById("formListIndx").valueAsNumber]
    
    li["name"] = document.getElementById("formListName").value
    li["description"] = document.getElementById("formListDesc").value
    await updateList(li)

    let nr = newListRow(li, document.getElementById("formListIndx").valueAsNumber+1)
    document.getElementById("rowListID-"+li["listID"]).outerHTML = nr

}

async function handleDeleteList(listIndx) {
    let li = LISTS[listIndx];
    if (ITEMS != null){
        for (let i = 0; i < ITEMS.length; i++) {
            const it = ITEMS[i];
            await deleteItem(it) // removes items from DB
        }
        clearItems()
    }
    let r = await deleteList(li); // Deletes the list
    console.log(r);
    clearLists();
    initialLoad();
}

async function handleCompleteList(listIndx) {
    let li = LISTS[listIndx]
    let ch = document.getElementById("listCheck-"+li["listID"]).checked
    li["status"] = ch ? 1 : 0
    await setListCompleted(li)
}

async function handleCompleteItem(itemIndx) {
    let it = ITEMS[itemIndx]
    let ch = document.getElementById("itCheck-"+it["itemID"]).checked
    it["status"] = ch ? 1 : 0
    await setitemCompleted(it)
}