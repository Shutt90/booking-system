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
let activeForm = qs('.active')

const options = [qs('#form-arrival'), qs('#form-depart')]

function calenderState() {

    // options.forEach((node, index) => {
        
    // })

    this.addEventListener('click', function() {
        this.classList.toggle('active');
    })

    if(this.classList.contains('active') && calender.style.display === 'flex') {
        this.classList.remove('active')
    } else if(!this.classList.contains('active')) {
        this.classList.add('active')
    }

    const activeFormNew = qs('.active')
    console.log(activeFormNew)

    if(activeFormNew.classList.contains('active')) {
        calender.style.display = 'flex'
    } else if(!activeFormNew === null){
        calender.style.display = 'none'
    }

}

function getValue(e) {
    const selected = e.target.textContent.trim()
    if(depart.classList.contains('active')) {
        depart.value = selected 
    } else if (arrival.classList.contains('active')) {
        arrival.value = selected
    }
}

depart.addEventListener('click', calenderState, this)
arrival.addEventListener('click', calenderState)

dateChildren.forEach(date => {
    date.addEventListener('click', getValue)
})