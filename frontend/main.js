function qs(node) {
    return document.querySelector(node)
}

const arrowLeft = qs('.arrow-left')
const arrowRight = qs('.arrow-right')
const calender = qs('.calender')
const date = qs('.calender-dates')
const day = qs('.calender-day')
const arrival = qs('#form-arrival')
const depart = qs('#form-depart')
const dateChildren = date.querySelectorAll('div')

function calenderState() {
    if(calender.style.display === 'none') {
        calender.style.display = 'flex'
    } else {
        calender.style.display = 'none'
    }
}

function getValue(e) {
    console.log(e.target.textContent.trim())
}

depart.addEventListener('click', calenderState)
arrival.addEventListener('click', calenderState)


dateChildren.forEach(date => {
    date.addEventListener('click', getValue)
})