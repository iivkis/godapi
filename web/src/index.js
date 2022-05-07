var ITEMS_LIST_WRAP = [];

window.onload = function () {
    ITEMS_LIST_WRAP = document.querySelectorAll('.items_list__wrap')
}

function toggleSubgroupVisibility(subgroup, event) {
    let el = document.getElementById("subgroup-" + subgroup + "-items")
    if (el.getAttribute('hidden') == null) {
        el.setAttribute('hidden', '')
        event.target.innerText = 'show';
    }
    else {
        el.removeAttribute('hidden')
        event.target.innerText = 'hide';
    }
}

function toggleItemVisibility(descID) {
    ITEMS_LIST_WRAP.forEach(el => {
        el.setAttribute('hidden', '')
    })
    document.getElementById(descID).removeAttribute('hidden')
}