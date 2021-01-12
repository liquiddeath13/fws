# fws
file server for any file format and size getting

# how it works
In <a href="https://github.com/liquiddeath13/fcs/">another repository</a> located client, which sends file with specified path to specified origin as run key, where is file server is on running state

# features
- using tcp conn
- many action types:
  - getting file from new origin by any run key will lead to dir creation with name like an remote IP and saving file in that folder
  - getting file from existed origin by run key "new" with existed name will lead to name changing like <{name} (N).{ext}>, where N is order num of files with the same name
  - getting file from existed origin by run key "append" will lead to addition of specified file
- ready to use solution: clone&build
