import Menu from './components/Menu'
import NavButton from './components/NavButton'
import { MenuType } from "./store/MenuStore"

function App() {

  return (
    <>
      <header className="
      max-w-full h-14 grow-0
      bg-gradient-to-r from-gray-900 to-zinc-950 backdrop-blur-sm
      flex pl-5
      font-normal text-2xl text-gray-500 hover:text-gray-300 transition duration-300
      ">
        <div className="content-center">
          A<span className='text-zinc-50'>Maze</span>ing <span className='text-zinc-50'>Generator</span>
        </div>
      </header>


      <main className="
      grow
      flex flex-col
      ">
      <nav className="
      w-full h-8
      flex flex-row space-x-5 pl-5
      ">

        <NavButton text="Generate" menuType={MenuType.Generate}/>
        <NavButton text="Solve" menuType={MenuType.Solve}/>

      </nav>

      <div className="
      grow
      flex flex-row
      ">

        <Menu/>

        <aside className='grow'>




        </aside>



      </div>



      </main>
      
    </>
  )
}

export default App
