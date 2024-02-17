const menu = document.querySelector('.menu')
const mobile_menu = document.querySelector('.toggle_menu')
const reposElement = document.getElementById('repos')
const load = document.getElementById('load')

let menu_toggle = false
let repo_amount = 3

const createRepoListElement = (element, name, url, description, topics) => {
  let list = document.createElement('li')
  list.classList.add('service')
  let a = document.createElement('a')
  a.href = url
  a.target = '_blank'
  let div = document.createElement('div')
  let h1 = document.createElement('h1')
  h1.innerHTML = name
  let arrow = document.createElement('i')
  arrow.classList.add('fa-solid', 'fa-arrow-right')
  let p = document.createElement('p')
  p.innerHTML = description
  let ul = document.createElement('ul')
  if(topics){
    topics.forEach(topic => {
      let li = document.createElement('li')
      li.classList.add('tag')
      let i = document.createElement('i')
      i.classList.add('fa-solid', 'fa-tag')
      li.innerHTML = topic
      li.appendChild(i)
      ul.appendChild(li)
    })
  }
  list.appendChild(a)
  a.appendChild(div)
  div.appendChild(h1)
  div.appendChild(arrow)
  a.appendChild(p)
  a.appendChild(ul)
  element.appendChild(list)
}

const resetRepoListElement = (element) => {
  Array.from(element.children).forEach(child => {
    element.removeChild(child)
  })
}


menu.addEventListener('click', () => {
  if (menu.classList.contains('fa-bars')) {
    menu.classList.remove('fa-bars')
    menu.classList.add('fa-x')
  } else {
    menu.classList.add('fa-bars')
    menu.classList.remove('fa-x')
  }
  menu.classList.add('active')
  menu_toggle ? mobile_menu.classList.remove('active') : mobile_menu.classList.add('active')
  menu_toggle = !menu_toggle
})

const getRepos = async () => {
  try {
    const response = await fetch('https://api.github.com/users/robertgouveia/repos', {
      method: 'GET',
      headers: {
        'Authorization': 'Bearer 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'
      }
    })
    if (response.ok) {
      return await response.json()
    } else {
      console.log('error retrieving data')
    }
  } catch (e) {
    console.log(e)
  }
}

const renderResponse = (repos) => {
  resetRepoListElement(reposElement)
  for(let i = 0; i < repo_amount; i++){
    createRepoListElement(reposElement, repos[i]['name'], repos[i]['html_url'], repos[i]['description'], repos[i]['topics'])
  }
}

load.addEventListener('click', async (event) => {
  event.preventDefault();
  repo_amount += 3
  const repos = await getRepos()
  if(repos) {
    renderResponse(repos)
  }
})

document.addEventListener('DOMContentLoaded', async () => {
  const repos = await getRepos()
  if(repos) {
    renderResponse(repos)
  }
})

