

function newListRow(list, listnum) {
    let row = `
<div class="panel centerContent" id="rowListID-${list["listID"]}" onclick="handleListClick(${list["listID"]}, ${listnum})">
    <input class="hide" type="number" name="listID" id="list-${list["listID"]}">
    <span class="largetext">${listnum}</span>
    <div class="flex-item-right centerContent">
        <span class="">${list["name"]}</span>
        <span class="">
            <label for="status">Completed</label>
            <input id="listCheck-${list["listID"]}" type="checkbox" name="status" onClick="handleCompleteList(${listnum-1})" ${list["status"] ? "checked" : ""}>
        </span>
        <span class="">${list["description"]}</span>
        <span class="">
            <button class="danger" onClick="handleDeleteList(${listnum-1})">Delete</button>
        </span>
    </div>
</div>
`;

    return row
}

function appendListRow(htmlRow) {
    let cont = document.getElementById("lists");
    cont.innerHTML = cont.innerHTML + htmlRow;
}
function clearLists(){
    document.getElementById("lists").innerHTML = ""
}


function newItemRow(item, itemIndx) {
    // let status = item["status"] ? "checked" : ""
    let row = `
<div class="panel centerContent" onclick="handleItemClick(${itemIndx})">
    <span>${item["name"]}</span>
    <span>${item["description"]}</span>
    <span class="">
        <label for="status">Completed</label>
        <input type="checkbox" name="status" id="itCheck-${item["itemID"]}" onclick="handleCompleteItem(${itemIndx})" ${item["status"] ? "checked" : ""}>
    </span>
    <span class="">
        <button class="danger" onClick="handleDeleteItem(${itemIndx})">Delete</button>
    </span>
</div>
`
    return row
}

function appendItemRow(htmlRow) {
    let cont = document.getElementById("itemView");
    cont.innerHTML = cont.innerHTML + htmlRow;
}
function clearItems(){
    document.getElementById("itemView").innerHTML = ""
}


function getListEleByID(listID) {
    document.getElementById("")
}


// Forms
function fillListForm(listIndex) {
    let l = LISTS[listIndex]
    document.getElementById("formListIndx").value = listIndex
    document.getElementById("formListID").value = l["listID"]
    document.getElementById("formListName").value = l["name"]
    document.getElementById("formListDesc").value = l["description"]
}

function fillItemForm(itemIndx) {
    let l = ITEMS[itemIndx]
    document.getElementById("formItemIndx").value = itemIndx
    document.getElementById("formItemID").value = l["itemID"]
    document.getElementById("formItemName").value = l["name"]
    document.getElementById("formItemDesc").value = l["description"]
}
