import Quill from 'quill'
import QuillTableBetter from 'quill-table-better'
import 'quill/dist/quill.snow.css'
import 'quill-table-better/dist/quill-table-better.css'

Quill.register('modules/table-better', QuillTableBetter, true)

export function initQuill(selector, tresc) { //tresc
  const toolbarOpt = [
    ['bold', 'italic', 'strike'],
    [{ header: [2, 3, 4, false] }],
    ['blockquote'],
    ['link', 'image', 'video'],
    [{ list: 'ordered' }, { list: 'bullet' }, { list: 'check' }],
    [{ script: 'sub' }, { script: 'super' }],
    [{ indent: '-1' }, { indent: '+1' }],
    [{ direction: 'rtl' }],
    [{ font: 'Sans' }],
    [{ align: [] }],
    ['table-better'],
    ['clean']
  ]

  const quill = new Quill(selector, {
    theme: 'snow',
    placeholder: 'Opis...',
    modules: {
      toolbar: toolbarOpt,
      'table-better': {
        toolbarTable: true,
        operationMenu: {
          items: {
            unmergeCells: true,
            deleteRow: true,
            deleteCol: true,
            deleteTable: true
          }
        }
      },
      keyboard: {
        bindings: QuillTableBetter.keyboardBindings
      }
    }
  })

  quill.on('text-change', () => {
    quill.root.querySelectorAll('img').forEach(img => {
      if (!img.hasAttribute('alt')) {
        const alt = prompt('alt=')
        img.setAttribute('alt', alt)
      }
    })
  })

  function preDBProcessing(tresc){
    const cleanHTML = tresc.replace(/^<temporary[^>]*>/, '<tbody>').replace(/<\/temporary>$/, '</tbody>')
    return cleanHTML
  }
  
  quill.root.innerHTML = preDBProcessing(tresc);
  return quill
}
