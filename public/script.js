async function fetchNames() {
  try {
    const response = await fetch('/api/generate-name', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    
    if (!response.ok) throw new Error('网络响应异常');
    
    const data = await response.text();
    const names = data.split('\n');
    
    // 更新DOM显示生成的名称
    const container = document.getElementById('names-container');
    container.innerHTML = names.map(name => `<div>${name}</div>`).join('');
  } catch (error) {
    console.error('请求失败:', error);
    alert('名称生成失败，请稍后重试');
  }
}

// 绑定按钮点击事件
document.getElementById('generate-btn').addEventListener('click', fetchNames);