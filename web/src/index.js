var LAST_OPEN_ITEM_ID = "";

function toggleSubgroupVisibility(subgroup, event) {
    let el = document.getElementById(`subgroup-${subgroup}-items`)
    if (el.getAttribute('hidden') == null) {
        el.setAttribute('hidden', '')
        event.target.innerText = 'show';
    }
    else {
        el.removeAttribute('hidden')
        event.target.innerText = 'hide';
    }
}

function toggleItemVisibility(itemID) {
    if (LAST_OPEN_ITEM_ID != itemID) {
        if (LAST_OPEN_ITEM_ID != "")
            document.getElementById(LAST_OPEN_ITEM_ID).setAttribute('hidden', '');
        document.getElementById(itemID).removeAttribute('hidden')
        LAST_OPEN_ITEM_ID = itemID;
    }
}