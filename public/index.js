$(document).ready(function() {
    $.ajax({
      url: '/getCollections',
      method: 'GET',
      success: function(response) {
        const $collectionList = $('<ul>'); // ul 요소는 유지 (하이퍼링크를 담을 컨테이너)
  
        $.each(response, function(index, collectionName) {
          const $link = $('<a>').attr('href', `/${collectionName}`).text(collectionName); // 하이퍼링크 생성
          const $listItem = $('<li>').append($link); // 하이퍼링크를 li 안에 추가
          $collectionList.append($listItem);
        });
  
        $('body').append($collectionList);
      },
    });
  });