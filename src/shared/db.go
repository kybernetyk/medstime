package main

// import (
//  "os"
//  "fmt"
//  //  "github.com/mikejs/gomongo/mongo"
//  "launchpad.net/gobson/bson"
//  "launchpad.net/mgo"
// )
// 
// const (
//  db_name      = "medstime"
//  col_accounts = "accounts"
// )
// 
// type MongoDB struct {
//  db   *mgo.Database
//  conn *mgo.Session
// }
// 
// type querymap map[string]interface{}
// 
// func NewMongoDB() *MongoDB {
//  d := &MongoDB{}
//  return d
// }
// 
// func (md *MongoDB) Connect() {
//  var err os.Error
// 
//  md.comm, err = mgo.Mongo("127.0.0.1")
//  if err != nil {
//      fmt.Println("Couldn't connect to mongo db @ localhost")
//      os.Exit(-1)
//      return
//  }
// 
//  md.db = md.conn.DB("medstime")
// }
// 
// func (md *MongoDB) Disconnect() {
//  md.conn.Close()
// }
// 
// func (md *MongoDB) Query(collection string, qryobj querymap, skip, limit int32) (docs []interface{}, err os.Error) {
//     query := md.db.C(collection).Find(qryobj)
// 
//  var query mongo.BSON
//  query, err = mongo.Marshal(qryobj)
//  if err != nil {
//      return
//  }
// 
//  var documents *mongo.Cursor
//  documents, err = md.db.GetCollection(collection).Query(query, skip, limit)
//  if err != nil {
//      return
//  }
// 
//  var doc mongo.BSON
//  for documents.HasMore() {
//      doc, err = documents.GetNext()
//      if err != nil {
//          return
//      }
//      docs = append(docs, doc)
//  }
//  return
// }
// 
// func (md *MongoDB) Count(collection string, qryobj querymap) int64 {
//  query := md.db.C(collection).Find(qryobj)
//  count, err := query.Count()
//  if err != nil {
//      return 0
//  }
// 
//  return count
// }
// 
// func (md *MongoDB) Update(collection string, object interface{}, qryobj querymap) bool {
//  err := md.db.C(collection).Update(qryobj, object)
//  if err != nil {
//      return false
//  }
//  return true
// }
// 
// 
// func (md *MongoDB) Insert(collection string, object interface{}) bool {
//  err = md.db.C(collection).Insert(object)
//  if err != nil {
//      return false
//  }
//  return true
// }
// 
// 
// // func (md *MongoDB) GetAccountForEmail(email string) (acc Account, err os.Error) {
// //  type q map[string]interface{}
// // 
// //  qry := q{
// //      "$query": q{"email": email},
// //  }
// // 
// //  var docs []mongo.BSON
// //  docs, err = md.getDocsForQuery(col_accounts, qry, 0, 1)
// //  if err != nil || len(docs) == 0 {
// //      err = os.NewError("Account not found.")
// //      return
// //  }
// // 
// //  err = mongo.Unmarshal(docs[0].Bytes(), &acc)
// //  return
// // }
// // 
// // func (md *MongoDB) GetAccountForAccountId(acc_id int64) (acc Account, err os.Error) {
// //  type q map[string]interface{}
// // 
// //  qry := q{
// //      "$query": q{"id": acc_id},
// //  }
// // 
// //  var docs []mongo.BSON
// //  docs, err = md.getDocsForQuery(col_accounts, qry, 0, 1)
// //  if err != nil || len(docs) == 0 {
// //      err = os.NewError("Account not found.")
// //      return
// //  }
// // 
// //  err = mongo.Unmarshal(docs[0].Bytes(), &acc)
// //  return
// // }
// // 
// // 
// // func (md *MongoDB) GetAccountCount() int64 {
// //  type q map[string]interface{}
// // 
// //  qry := q{
// //      "$query": q{},
// //  }
// // 
// //  return md.getCountForQuery(col_accounts, qry)
// // }
// // 
// // func (md *MongoDB) StoreAccount(account Account) (acc_id int64, err os.Error) {
// // 
// //  //create new acc
// //  if account.Id == 0 {
// //      qry, _ := mongo.Marshal(map[string]string{})
// //      count, _ := md.db.GetCollection(col_accounts).Count(qry)
// //      count++
// // 
// //      acc_id = count
// //      account.Id = acc_id
// //      doc, _ := mongo.Marshal(account)
// //      err = md.db.GetCollection(col_accounts).Insert(doc)
// //      return
// //  } else { //update post
// //      type q map[string]interface{}
// //      m := q{"id": account.Id}
// // 
// //      var query mongo.BSON
// //      query, err = mongo.Marshal(m)
// //      if err != nil {
// //          return
// //      }
// // 
// //      doc, _ := mongo.Marshal(account)
// //      err = md.db.GetCollection(col_accounts).Update(query, doc)
// //      acc_id = account.Id
// //  }
// // 
// //  return
// // }
// 
// 
// // func (md *MongoDB) Query(collection string, qryobj querymap, skip, limit int32) (docs []mongo.BSON, err os.Error) {
// //  var query mongo.BSON
// //  query, err = mongo.Marshal(qryobj)
// //  if err != nil {
// //      return
// //  }
// // 
// //  var documents *mongo.Cursor
// //  documents, err = md.db.GetCollection(collection).Query(query, skip, limit)
// //  if err != nil {
// //      return
// //  }
// // 
// //  var doc mongo.BSON
// //  for documents.HasMore() {
// //      doc, err = documents.GetNext()
// //      if err != nil {
// //          return
// //      }
// //      docs = append(docs, doc)
// //  }
// //  return
// // }
// // 
// // func (md *MongoDB) Count(collection string, qryobj querymap) int64 {
// //  var query mongo.BSON
// //  query, err := mongo.Marshal(qryobj)
// //  if err != nil {
// //      return 0
// //  }
// // 
// //  count, err := md.db.GetCollection(collection).Count(query)
// //  if err != nil {
// //      return 0
// //  }
// // 
// //  return count
// // }
// // 
// // func (md *MongoDB) Update(collection string, object interface{}, qryobj querymap) bool {
// //     query, err := mongo.Marshal(qryobj)
// //     if err != nil {
// //      return false
// //     }
// // 
// //     doc, _ := mongo.Marshal(object)
// //     err = md.db.GetCollection(collection).Update(query, doc)
// //     if err != nil {
// //         return false
// //     }
// //     return true
// // }
// // 
// // 
// // func (md *MongoDB) Insert(collection string, object interface{}) bool {
// //     doc, _ := mongo.Marshal(object)
// //     err := md.db.GetCollection(collection).Insert(doc)
// //     if err != nil {
// //         return false
// //     }
// //     return true
// // 
// // }
// 
// /*
// //warning: it will marhsall the comments list - so we need to change this
// //if we enable updating/editing posts
// func (md *MongoDB) StorePost(post *BlogPost) (id int64, err os.Error) {
//  md.postmu.Lock()
//  defer md.postmu.Unlock()
// 
//  //create new post
//  if post.Id == 0 {
//      qry, _ := mongo.Marshal(map[string]string{})
//      count, _ := md.posts.Count(qry)
//      count++
// 
//      id = count
//      post.Id = count
//      doc, _ := mongo.Marshal(*post)
//      err = md.posts.Insert(doc)
//      return
//  } else { //update post
//      type q map[string]interface{}
//      m := q{"id": post.Id}
// 
//      var query mongo.BSON
//      query, err = mongo.Marshal(m)
//      if err != nil {
//          return
//      }
// 
//         doc, _ := mongo.Marshal(*post)
//      err = md.posts.Update(query, doc)
//  }
// 
//  return
// }
// 
// func (md *MongoDB) getPostsForQuery(qryobj interface{}, skip, limit int32) (posts []BlogPost, err os.Error) {
//  md.postmu.Lock()
//  defer md.postmu.Unlock()
// 
//  var query mongo.BSON
//  query, err = mongo.Marshal(qryobj)
//  if err != nil {
//      return
//  }
// 
//  // count, _ := md.posts.Count(query)
//  // if count == 0 {
//  //  err = os.NewError("COUNT 0 Post Not Found")
//  //  return
//  // }
// 
//  var docs *mongo.Cursor
//  //docs, err = md.posts.FindAll(query)
//  docs, err = md.posts.Query(query, skip, limit)
//  if err != nil {
//      return
//  }
// 
//  var doc mongo.BSON
//  for docs.HasMore() {
//      doc, err = docs.GetNext()
//      if err != nil {
//          return
//      }
//      var post BlogPost
//      err = mongo.Unmarshal(doc.Bytes(), &post)
//      if err != nil {
//          return
//      }
//      posts = append(posts, post)
//  }
//  // if len(posts) == 0 {
//  //     err = os.NewError("no posts found")
//  // }
//  return
// }
// 
// 
// func (md *MongoDB) GetPost(post_id int64) (post BlogPost, err os.Error) {
//  type q map[string]interface{}
//  m := q{"id": post_id}
// 
//  //find sort example
//  // m := q{
//  //     "$query": q{"id": q{"$gte" : post_id}},
//  //     "$orderby": q{"timestamp": -1},
//  // }
// 
//  var posts []BlogPost
//  posts, err = md.getPostsForQuery(m, 0, 1)
//  if err != nil || len(posts) == 0 {
//      err = os.NewError("Post not found.")
//      return
//  }
// 
//  post = posts[0]
//  return
// }
// 
// //returns posts for a certain date
// func (md *MongoDB) GetPostsForDate(date time.Time) (posts []BlogPost, err os.Error) {
//  date.Hour = 0
//  date.Minute = 0
//  date.Second = 0
// 
//  start := date.Seconds()
//  end := start + (24 * 60 * 60)
// 
//  return md.GetPostsForTimespan(start, end, -1)
// }
// 
// //returns posts for a certain month
// func (md *MongoDB) GetPostsForMonth(date time.Time) (posts []BlogPost, err os.Error) {
//  date.Hour = 0
//  date.Minute = 0
//  date.Second = 0
//  date.Day = 1
// 
//  next_month := date
//  next_month.Month++
//  if next_month.Month > 12 {
//      next_month.Month = 1
//      next_month.Year++
//  }
// 
//  start := date.Seconds()
//  end := next_month.Seconds()
// 
//  return md.GetPostsForTimespan(start, end, -1)
// }
// 
// 
// func (md *MongoDB) GetPostsForTimespan(start_timestamp, end_timestamp , order int64) (posts []BlogPost, err os.Error) {
//  type q map[string]interface{}
// 
//  //  m := q{"id": post_id}
// 
//  m := q{
//      "$query":   q{"timestamp": q{"$gte": start_timestamp, "$lt": end_timestamp}},
//      "$orderby": q{"timestamp": order},
//  }
// 
//  posts, err = md.getPostsForQuery(m, 0, 0)
//  if err != nil || len(posts) == 0 {
//      err = os.NewError("Posts not found.")
//      return
//  }
// 
//  return
// }
// 
// func (md *MongoDB) GetLastNPosts(num_to_get int32) (posts []BlogPost, err os.Error) {
//  type q map[string]interface{}
//  m := q{
//      "$query":   q{},
//      "$orderby": q{"timestamp": -1},
//  }
// 
//  //var posts []BlogPost
//  posts, err = md.getPostsForQuery(m, 0, num_to_get)
//  if err != nil || len(posts) == 0 {
//      err = os.NewError("Posts not found.")
//      return
//  }
// 
//  return
// }
// 
// func (md *MongoDB) StoreComment(comment *PostComment) (id int64, err os.Error) {
//  md.commentmu.Lock()
//  defer md.commentmu.Unlock()
// 
//  //check if post with that id exists
//  _, err = md.GetPost(comment.PostId)
//  if err != nil {
//      //err = os.NewError("Post doesn't exist :]")
//      return
//  }
//  content := comment.Content
//  //content = strings.Replace(content, "<", "(", -1)
// //   content = strings.Replace(content, ">", ")", -1)
// 
//  //author := html.EscapeString(comment.Author)
//  author := comment.Author
// //   author = strings.Replace(author, "<", "(", -1)
//  //author = strings.Replace(author, ">", ")", -1)
// 
//  comment.Author = author//html.EscapeString(comment.Author)
//  comment.Content = content//html.EscapeString(comment.Content)
// 
//  qry, _ := mongo.Marshal(map[string]string{})
//  count, _ := md.comments.Count(qry)
//  count++
//  id = count
//  comment.Id = count
//  doc, _ := mongo.Marshal(*comment)
//  fmt.Println(doc)
// 
//  md.comments.Insert(doc)
// 
//  return
// }
// 
// //get comments belonging to a post
// func (md *MongoDB) GetComments(post_id int64) (comments []PostComment, err os.Error) {
//  md.commentmu.Lock()
//  defer md.commentmu.Unlock()
// 
//  //m := map[string]int64{"postid": post_id}
//  type q map[string]interface{}
// 
//      m := q{
//      "$query":   q{"postid": post_id},
//      "$orderby": q{"timestamp": 1},
//  }
// 
// 
//  var query mongo.BSON
//  query, err = mongo.Marshal(m)
//  if err != nil {
//      return
//  }
// 
//  var docs *mongo.Cursor
//  docs, err = md.comments.FindAll(query)
//  if err != nil {
//      return
//  }
// 
//  var doc mongo.BSON
// 
//  for docs.HasMore() {
//      doc, err = docs.GetNext()
//      if err != nil {
//          return
//      }
//      var comment PostComment
//      err = mongo.Unmarshal(doc.Bytes(), &comment)
//      if err != nil {
//          return
//      }
//      comments = append(comments, comment)
//  }
//  return
// }*/
