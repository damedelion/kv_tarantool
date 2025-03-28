box.cfg {
    listen = 3301
}

-- При поднятии БД создаем спейсы и индексы
box.once('init', function()
  kv = box.schema.space.create('kv')
  kv:format({
    { name = 'key', type = 'unsigned' },
    { name = 'value', type = 'string' },
  })
-- Create indexes --
  kv:create_index('primary', { parts = { 'key' } })

  print('Hello, world!')
end)

-- Даем доступ для юзера guest для подключения по guest пользователю
box.once('access:v1', function()
  box.schema.user.grant('guest', 'read,write', 'space', 'kv')
end)